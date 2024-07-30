package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}

		fmt.Fprintf(w, "Welcome to the home page!")
	})

	s := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("listening on port %s", port)
	log.Fatal(s.ListenAndServe())
}
