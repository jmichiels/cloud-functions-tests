package domain

import "github.com/pkg/errors"

type Account struct {
	Id       UniqueId
	ClientId UniqueId
	Balance  Money
}

// Validate ensures the Account is valid.
func (account *Account) Validate() error {
	if err := account.Id.Validate(); err != nil {
		return errors.Wrap(err, "invalid id")
	}
	if err := account.ClientId.Validate(); err != nil {
		return errors.Wrap(err, "invalid client id")
	}
	if account.Balance.IsNegative() {
		return errors.New("negative balance")
	}
	return nil
}

// Transfer transfers money from this Account to another.
func (account *Account) Transfer(amount Money, to *Account) error {
	remainingBalance := account.Balance.Subtract(amount)
	if remainingBalance.IsNegative() {
		return errors.New("insufficient funds")
	}
	account.Balance = remainingBalance
	to.Balance = to.Balance.Add(amount)
	return nil
}
