package server

import (
	"io/fs"
	"log"
	"net/http"
)

func (s *Server) MakeHandler() {
	mux := s.mux

	mux.HandleFunc("/", s.homeHandler)

	mux.Handle("/public/", s.staticHandler())
	mux.Handle("/favicon.ico", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("on get favicon.ico")

		serveFileFS(w, r, s.staticFS, "favicon/favicon.ico")
	}))

	mux.HandleFunc("/about", s.staticPageHandler("about", "About Page"))

	mux.HandleFunc("GET /accounts", s.getAccountsHandler)
	mux.HandleFunc("GET /test", s.testHandler)
}

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	s.servePage(w, "home", BasePage{
		HTMLTitle: "Home Page",
	})
}

func (s *Server) staticHandler() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.FS(s.staticFS)))
}

func serveFileFS(w http.ResponseWriter, r *http.Request, fsys fs.FS, name string) {
	fs := http.FileServer(http.FS(fsys))
	r.URL.Path = name
	fs.ServeHTTP(w, r)
}
