package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegram01/wingram/model"
)

type AccountsPage struct {
	BasePage
	Accounts []model.Account
}

func (s *Server) getAccountsHandler(w http.ResponseWriter, r *http.Request) {
	accounts, err := s.db.GetAccounts()
	if err != nil {
		log.Printf("getAccountsHandler error: %v", err)
		// just for test if error, in product we must change it
		http.Error(w, fmt.Sprintf("getAccountsHandler error: %v", err), http.StatusInternalServerError)
		return
	}

	// fmt.Println("Accounts is ", accounts)

	s.servePage(w, "account", AccountsPage{
		BasePage: s.newBasePage(r, "Accounts Page"),
		Accounts: accounts,
	})
}