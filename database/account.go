package database

import (
	"fmt"

	"github.com/codegram01/wingram/model"
)

func (db *Db)InsertAccount(accInp *model.Account) (model.Account, error) {
	var acc model.Account

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


func (db *Db)GetAccounts() ([]model.Account, error) {
	var accounts []model.Account

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
		var acc model.Account
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