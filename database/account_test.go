package database

import (
	"log"
	"testing"

	"github.com/codegram01/wingram/config"
	"github.com/codegram01/wingram/model"
)

func TestInsertAccount(t *testing.T) {
	cfg := config.Init()
	db, err := Connect(cfg)

	if err != nil {
		log.Fatal(err)
	}

	account, err := db.InsertAccount(&model.Account{
		Name: "alex",
		Email: "alex@gmail.com",
		Bio: "From England",
	})
	if err != nil {
		log.Fatal(err)
	}

	t.Log(account)
}

func TestGetAccounts(t *testing.T) {
	cfg := config.Init()
	db, err := Connect(cfg)

	if err != nil {
		log.Fatal(err)
	}

	accounts, err := db.GetAccounts()
	if err != nil {
		log.Fatal(err)
	}

	t.Log(accounts)
}