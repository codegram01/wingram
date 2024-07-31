package server

import (
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/codegram01/wingram/database"
	"github.com/codegram01/wingram/server/templates"
	"github.com/google/safehtml/template"
)

type Server struct {
	templates map[string]*template.Template
	mux       *http.ServeMux
	staticFS  fs.FS

	db *database.Db
}

// Config when init server
type ServerCfg struct {
	Port string
	Db *database.Db
}

func Init(cfg *ServerCfg) {
	ts, err := templates.ParsePageTemplates()
	if err != nil {
		log.Fatalf("error parsing templates: %v", err)
	}

	mux := http.NewServeMux()
	server := &Server{
		mux:       mux,
		templates: ts,
		staticFS:  os.DirFS("static/public"),
	}
	server.MakeHandler()

	s := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: mux,
	}

	log.Printf("Server running on Port: %s", cfg.Port)
	log.Fatal(s.ListenAndServe())
}
