package main

import (
	"log"
	"testing"

	"github.com/codegram01/wingram/config"
	"github.com/codegram01/wingram/database"
	// "github.com/codegram01/wingram/server"
)

func TestMain(t *testing.T) {
	cfg := config.Init()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	accounts, err := db.GetAccounts()
	if err != nil {
		log.Fatal(err)
	}

	t.Log(accounts)

}