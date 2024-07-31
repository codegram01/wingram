package database

import (
	"fmt"
)

type Account struct {
	Id int
	Name string
	Email string
	Bio string
}

func (db *Db)InsertAccount(accInp *Account) (Account, error) {
	var acc Account

	queryStr := `
		INSERT INTO account 
			(name, email, bio)
		VALUES 
			($1, $2, $3) 
		RETURNING id, name, email, bio
	`
	row := db.Con.QueryRow(queryStr, accInp.Name, accInp.Email, accInp.Bio)

	err := row.Scan(
		&acc.Id,
		&acc.Name,
		&acc.Email,
		&acc.Bio,
	)

	if err != nil {
		return acc, err
	}

	return acc, nil
}


func (db *Db)GetAccounts() ([]Account, error) {
	var accounts []Account

	queryStr := `
		SELECT 
			id, name, email, bio
		FROM 
			account
	`
	rows, err := db.Con.Query(queryStr)
	if err != nil {
		return nil, fmt.Errorf("GetAccounts: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var acc Account
		if err := rows.Scan(&acc.Id, &acc.Name, &acc.Email, &acc.Bio); err != nil {
			return nil, fmt.Errorf("GetAccounts: %v", err)
		}
		accounts = append(accounts, acc)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAccounts: %v", err)
	}

	return accounts, nil
}