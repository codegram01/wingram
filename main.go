package main

import (
	"log"

	"github.com/codegram01/wingram/config"
	"github.com/codegram01/wingram/database"
	"github.com/codegram01/wingram/server"
)

func main() {
	cfg := config.Init()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	// db, cleanup := database.ConnectGCloud()
  	// defer cleanup()

	sCfg := &server.ServerCfg{
		Cfg: cfg,
		Db: db,
	}
	server.Init(sCfg)
}
