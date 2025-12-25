package local

import (
	"errors"
	"math"
	"sort"
	"sync"

	"splitdim/pkg/api"
)

// localDB a simple implementation of the DataLayer API.
type localDB struct { //This is private due to the small first letter
	// accounts maintains the balance for each user name
	accounts map[string]int
	// The read-write mutex makes sure concurrent access is safe.
	mu sync.RWMutex
}

// NewDataLayer creates a new database of accounts.
func NewDataLayer() api.DataLayer {
	return &localDB{accounts: make(map[string]int)}
	//Idiomatic Go: data structure definition is private, so return interface as a pointer to the struct instead of the actual struct itself
}

func (db *localDB) Transfer(t api.Transfer) error {
	//Check if the sender and the receiver in the transfer are different and return an appropriate error if not
	if t.Sender == t.Receiver {
		return errors.New("Sender and Receiver cannot be the same in a transfer")
	}
	db.mu.Lock()
	defer db.mu.Unlock()
	//Perform the actual transaction

	//Do I care wether the value is zero or it did not exist before?
	//probably not, I have to initialize them anyway
	db.accounts[t.Sender] += t.Amount
	db.accounts[t.Receiver] -= t.Amount

	return nil
}

func (db *localDB) AccountList() ([]api.Account, error) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	ret := []api.Account{}
	for name, balance := range db.accounts {
		ret = append(ret, api.Account{Holder: name, Balance: balance})
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].Holder < ret[j].Holder
	})
	return ret, nil
}

func copyMap(original map[string]int) map[string]int {
	copy := make(map[string]int)
	for key, value := range original {
		copy[key] = value
	}
	return copy
}

func (db *localDB) Clear() ([]api.Transfer, error) {
	db.mu.RLock()
	//Database consistency check: summing up the balances
	var sum int
	for _, value := range db.accounts {
		sum += value
	}
	if sum != 0 {
		return nil, errors.New("Database is inconsistent")
	}

	tempAcc := copyMap(db.accounts)
	db.mu.RUnlock()
	transfers := []api.Transfer{}
	for sender, balance := range tempAcc {
		if balance < 0 {
			for receiver, receiverBalance := range tempAcc {
				if receiverBalance > 0 {
					float_balance := math.Abs(float64(balance))
					float_receiverBalance := math.Abs(float64(receiverBalance))
					//compute the minimum of the balances of the sender and receiver and store it in transferAmount
					transferAmount := math.Min(float_balance, float_receiverBalance)
					transfers = append(transfers, api.Transfer{Sender: sender, Receiver: receiver, Amount: int(transferAmount)})
					tempAcc[sender] += int(transferAmount)
					tempAcc[receiver] -= int(transferAmount)
					if tempAcc[sender] == 0 {
						break
					}
				}
			}
		}
	}

	return transfers, nil
}

func (db *localDB) Reset() error {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.accounts = make(map[string]int)
	return nil
}
