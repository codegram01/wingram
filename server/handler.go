package server

import (
	"io/fs"
	"log"
	"net/http"
)

func (s *Server) MakeHandler() {
	mux := s.mux

	mux.HandleFunc("/", s.homeHandler)
	mux.HandleFunc("/about", s.staticPageHandler("about", "About Page"))

	mux.Handle("/public/", s.staticHandler())
	mux.Handle("/favicon.ico", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("on get favicon.ico")

		serveFileFS(w, r, s.staticFS, "favicon/favicon.ico")
	}))
}

func (s *Server) homeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
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
