package server

import (
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/codegram01/wingram/config"
	"github.com/codegram01/wingram/database"
	"github.com/codegram01/wingram/server/templates"
	"github.com/google/safehtml/template"
)

type Server struct {
	mode string

	templates map[string]*template.Template
	mux       *http.ServeMux
	staticFS  fs.FS

	db *database.Db
	
}

// Config when init server
type ServerCfg struct {
	Cfg *config.Config
	Db *database.Db
}

func Init(scfg *ServerCfg) {
	if scfg.Db == nil {
		log.Fatal("Server need database to run")
	}

	ts, err := templates.ParsePageTemplates()
	if err != nil {
		log.Fatalf("error parsing templates: %v", err)
	}

	mux := http.NewServeMux()
	server := &Server{
		mode: scfg.Cfg.Mode,
		mux:       mux,
		templates: ts,
		staticFS:  os.DirFS("static/public"),
		db: scfg.Db,
		
	}
	server.MakeHandler()

	s := &http.Server{
		Addr:    ":" + scfg.Cfg.Port,
		Handler: mux,
	}

	log.Printf("Server running on Port: %s", scfg.Cfg.Port)
	log.Fatal(s.ListenAndServe())
}
