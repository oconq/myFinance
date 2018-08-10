package models

import (
	"awesomeProject/config"
	"log"
	"fmt"
)

type Account struct {
	Id 			int			`json:"id"`
	Name		string		`json:"name"`
}

type Accounts []Account


func FindAccountById(id int) *Account {
	var account Account

	row := config.Db().QueryRow("SELECT id, name FROM accounts WHERE id = $1", id)
	err := row.Scan(&account.Id, &account.Name)
	if err != nil {
		log.Fatal(err)
	}

	return &account
}

func AllAccounts() *Accounts {
	var accounts Accounts

	fmt.Println("Entering AllAccounts...")

	rows, err := config.Db().Query("SELECT id, name FROM accounts")
	if err != nil {
		log.Fatal(err)
	}

	// Close rows afeter all readed
	defer rows.Close()

	for rows.Next() {
		var a Account

		err := rows.Scan(&a.Id, &a.Name)
		if err != nil {
			log.Fatal(err)
		}

		accounts = append(accounts, a)
	}

	return &accounts
}

