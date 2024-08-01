package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/codegram01/wingram/config"
	_ "github.com/lib/pq"

	"cloud.google.com/go/cloudsqlconn"
	"cloud.google.com/go/cloudsqlconn/postgres/pgxv4"
)

type Db struct {
	Con *sql.DB // Database connection
}

func Connect(cfg *config.Config) (*Db, error) {
	con, err := sql.Open("postgres", cfg.GetDbConStr())
	if err != nil {
		return nil, err
	}

	pingErr := con.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	log.Println("DB Connected!")

	db := &Db{
		Con: con,
	}

	return db, nil
}

// getDB creates a connection to the database
// based on environment variables.
func ConnectGCloud() (*Db, func() error) {
	cleanup, err := pgxv4.RegisterDriver("cloudsql-postgres", cloudsqlconn.WithIAMAuthN())
	if err != nil {
	  log.Fatalf("Error on pgxv4.RegisterDriver: %v", err)
	}
  
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable", "inner-legacy-430517-g0:us-central1:wingram", "wingram@inner-legacy-430517-g0.iam", "wingram")
	con, err := sql.Open("cloudsql-postgres", dsn)
	if err != nil {
	  log.Fatalf("Error on sql.Open: %v", err)
	}

	db := &Db{
		Con: con,
	}

	log.Println("Connect to GSQL success")
  
	return db, cleanup
  }
  