package main

import (
	"bankcore"
	"testing"
)

func TestAccount(t *testing.T) {
	account := bankcore.Account{
		Customer: bankcore.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	if account.Name == "" {
		t.Error("can't create an Account object")
	}
}

func TestDeposit(t *testing.T) {
	account := bankcore.Account{
		Customer: bankcore.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	account.Deposit(10)
	if account.Balance != 10 {
		t.Errorf("balance is not being updated after a deposit")
	}
}

func TestWithdraw(t *testing.T) {
	account := bankcore.Account{Customer: bankcore.Customer{
		Name:    "John",
		Address: "Los Angeles, California",
		Phone:   "(213) 555 0147",
	},
		Number:  1001,
		Balance: 0,
	}
	account.Deposit(10)
	account.Withdraw(10)
	if account.Balance != 0 {
		t.Errorf("balance is not being updated after a withdraw")
	}
}

func TestStatement(t *testing.T) {
	account := bankcore.Account{
		Customer: bankcore.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	account.Deposit(100)
	statement := account.Statement()
	if statement != "1001 - John - 100" {
		t.Errorf("statement doesn't have the proper format")
	}
}
