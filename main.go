package main

import (
	"bankcore"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var accounts = map[float64]*bankcore.Account{}

func statement(w http.ResponseWriter, r *http.Request) {
	numberqs := r.URL.Query().Get("number")
	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}
	number, err := strconv.ParseFloat(numberqs, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			fmt.Fprint(w, account.Statement())
		}
	}
}

func deposit(w http.ResponseWriter, r *http.Request) {
	numberqs := r.URL.Query().Get("number")
	amountqs := r.URL.Query().Get("amount")
	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}
	number, err := strconv.ParseFloat(numberqs, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid account number!")
		return
	}
	amount, err := strconv.ParseFloat(amountqs, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
		return
	}
	acount, ok := accounts[number]
	if !ok {
		fmt.Fprintf(w, "Account with number %v can't be found!", number)
	} else {
		err := acount.Deposit(amount)
		if err != nil {
			fmt.Fprintf(w, "%v", err)
		} else {
			fmt.Fprint(w, acount.Statement())
		}
	}
}

func withdraw(w http.ResponseWriter, r *http.Request) {
	numberqs := r.URL.Query().Get("number")
	amountqs := r.URL.Query().Get("amount")
	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}
	number, err := strconv.ParseFloat(numberqs, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid account number!")
		return
	}
	amount, err := strconv.ParseFloat(amountqs, 64)
	if err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
		return
	}
	acount, ok := accounts[number]
	if !ok {
		fmt.Fprintf(w, "Account with number %v can't be found!", number)
		return
	}
	err = acount.Withdraw((amount))
	if err != nil {
		fmt.Fprintf(w, "%v", err)
	} else {
		fmt.Fprint(w, acount.Statement())
	}
}

func main() {
	accounts[1001] = &bankcore.Account{
		Customer: bankcore.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}
	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
