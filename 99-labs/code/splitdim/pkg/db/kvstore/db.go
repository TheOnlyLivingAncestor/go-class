package kvstore

import (
	"errors"
	"fmt"
	clientapi "kvstore/pkg/api"
	"kvstore/pkg/client"
	"resilient"
	"splitdim/pkg/api"
	"splitdim/pkg/clear"
	"strconv"
	"time"
)

type kvstore struct {
	client.Client
}

func NewDataLayer(kvstoreAddr string) api.DataLayer {
	return &kvstore{Client: client.NewClient(kvstoreAddr)}
}

func (db *kvstore) setBalance(user string, amount int) error {
	vv, err := db.Get(user)
	if err != nil {
		return err
	}
	balance, _ := strconv.Atoi(vv.Value)
	vv.Value = fmt.Sprintf("%d", balance+amount)
	vkv := clientapi.VersionedKeyValue{Key: user, VersionedValue: vv}
	err = db.Put(vkv)
	if err != nil {
		return err
	}
	return nil
}

func (db *kvstore) setBalanceForUser(user string, amount int) resilient.Closure {
	return func() error { return db.setBalance(user, amount) }
}

func (db *kvstore) Transfer(t api.Transfer) error {
	if t.Sender == t.Receiver || t.Sender == "" || t.Receiver == "" {
		return errors.New("Invalid transfer")
	}
	var defaultBackoff = resilient.Backoff{
		Base:      150 * time.Millisecond,
		NumTrials: 6,
		Cap:       2 * time.Second,
		Jitter:    3,
	}

	//Transfer for sender
	var senderBalanceClosure = db.setBalanceForUser(t.Sender, t.Amount)
	var decoratedSenderBalanceClosure = resilient.WithRetry(senderBalanceClosure, defaultBackoff)
	err := decoratedSenderBalanceClosure()
	if err != nil {
		return err
	}

	//Transfer for receiver
	var receiverBalanceClosure = db.setBalanceForUser(t.Receiver, -t.Amount)
	var decoratedReceiverBalanceClosure = resilient.WithRetry(receiverBalanceClosure, defaultBackoff)
	err = decoratedReceiverBalanceClosure()
	if err != nil {
		//Undo the first operation? t.Sender, -t.Amount?
		//More agressive retry policy?
		var aggressiveBackoff = resilient.Backoff{
			Base:      100 * time.Millisecond,
			NumTrials: 10,
			Cap:       2 * time.Second,
			Jitter:    2,
		}
		var senderUndoClosure = db.setBalanceForUser(t.Sender, -t.Amount)
		var decoratedSenderUndoClosure = resilient.WithRetry(senderUndoClosure, aggressiveBackoff)
		err = decoratedSenderUndoClosure()
		if err != nil {
			return err
		}
	}
	return nil
}

func (db *kvstore) AccountList() ([]api.Account, error) {
	accounts, err := db.List()
	if err != nil {
		return []api.Account{}, err
	}
	ret := []api.Account{}
	for _, account := range accounts {
		balance, _ := strconv.Atoi(account.Value)
		ret = append(ret, api.Account{Holder: account.Key, Balance: balance})
	}
	return ret, nil
}

func (db *kvstore) Clear() ([]api.Transfer, error) {
	//List the account database
	accounts, err := db.List()
	if err != nil {
		return []api.Transfer{}, err
	}
	//accounts is []api.VersionedKeyValue{}, create a map[string]int from them
	accountsMap := make(map[string]int)
	for _, account := range accounts {
		vv, _ := strconv.Atoi(account.Value)
		accountsMap[account.Key] = vv
	}
	transfers, err := clear.Clear(accountsMap)
	if err != nil {
		return []api.Transfer{}, err
	}
	return transfers, nil
}

// Reset sets all balances to zero.
func (db *kvstore) Reset() error {
	return db.Client.Reset()
}
