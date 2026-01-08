package clear

import (
	"math"
	"splitdim/pkg/api"
)

// Clear clears the debts for the accounts given as argument. Meanwhile, it updates "accounts", so always pass a copy to this function.
func Clear(accounts map[string]int) ([]api.Transfer, error) {
	transfers := []api.Transfer{}
	for sender, balance := range accounts {
		if balance < 0 {
			for receiver, receiverBalance := range accounts {
				if receiverBalance > 0 {
					float_balance := math.Abs(float64(balance))
					float_receiverBalance := math.Abs(float64(receiverBalance))
					//compute the minimum of the balances of the sender and receiver and store it in transferAmount
					transferAmount := math.Min(float_balance, float_receiverBalance)
					transfers = append(transfers, api.Transfer{Sender: sender, Receiver: receiver, Amount: int(transferAmount)})
					accounts[sender] += int(transferAmount)
					accounts[receiver] -= int(transferAmount)
					if accounts[sender] == 0 {
						break
					}
				}
			}
		}
	}
	return transfers, nil
}
