package server

import "net/http"

type TestPage struct {
	BasePage
	Mode string
}

func (s *Server) testHandler(w http.ResponseWriter, r *http.Request) {
	s.servePage(w, "test", TestPage{
		BasePage: s.newBasePage(r, "Test Page"),
		Mode: s.mode,
	})
}