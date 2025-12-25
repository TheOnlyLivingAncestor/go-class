package api

// Transfer represents a transaction.
type Transfer struct {
	// The deptor name.
	Sender string `json:"sender"`
	// The creditor name.
	Receiver string `json:"receiver"`
	// The amount transferred in the transaction.
	Amount int `json:"amount"`
}

// Account represents the balance of a user.
type Account struct {
	// The name of the account holder.
	Holder string `json:"holder"`
	// Current balance.
	Balance int `json:"balance"`
}

// DataLayer is an API for manipulating a balance sheet.
type DataLayer interface {
	// Transfer will process a transfer.
	Transfer(t Transfer) error
	// AccountList returns the current account of each user.
	AccountList() ([]Account, error)
	// Clear returns the list of transfers to clear all debts.
	Clear() ([]Transfer, error)
	// Reset sets all balances to zero.
	Reset() error
}
