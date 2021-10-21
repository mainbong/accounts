package accounts

import "errors"

var errNoMoney = errors.New("you have no money")

type Account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *Account {
	return &Account{owner: owner, balance: 1000}
}

// Deposit (입금)
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Withdraw (출금)
func (a *Account) Withdraw(amount int) error {
	// 마이너스가 안될때만 출금
	if a.balance >= amount {
		a.balance -= amount
		return nil
	} else {
		return errNoMoney
	}
}

// Owner
func (a Account) Owner() string {
	return a.owner
}

// Balance
func (a Account) Balance() int {
	return a.balance
}
