package database

import (
	"fmt"
	"log"
	"testing"

	"github.com/codegram01/wingram/config"
)

func TestConnectDb(t *testing.T) {
	cfg := config.Init()

	c, err := Connect(cfg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(c)
}
