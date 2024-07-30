package server

import (
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/codegram01/wingram/server/templates"
	"github.com/google/safehtml/template"
)

type Server struct {
	templates  map[string]*template.Template
	mux *http.ServeMux
	staticFS fs.FS
}

func Init() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	ts, err := templates.ParsePageTemplates()
	if err != nil {
		log.Fatalf("error parsing templates: %v", err)
	}

	mux := http.NewServeMux()
	server := &Server{
		mux: mux,
		templates: ts,
		staticFS: os.DirFS("static/public"),
	}
	server.MakeHandler()

	s := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("listening on port %s", port)
	log.Fatal(s.ListenAndServe())
}
