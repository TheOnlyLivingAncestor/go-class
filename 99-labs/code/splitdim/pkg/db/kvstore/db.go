package kvstore

import (
	"errors"
	"fmt"
	clientapi "kvstore/pkg/api"
	"kvstore/pkg/client"
	"strconv"

	"splitdim/pkg/api"
	"splitdim/pkg/clear"
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

func (db *kvstore) Transfer(t api.Transfer) error {
	if t.Sender == t.Receiver || t.Sender == "" || t.Receiver == "" {
		return errors.New("Invalid transfer")
	}
	for {
		err := db.setBalance(t.Receiver, -t.Amount)
		if err == nil {
			break
		}
	}
	for {
		err := db.setBalance(t.Sender, t.Amount)
		if err == nil {
			break
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
