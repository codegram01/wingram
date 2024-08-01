package server

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/google/safehtml"
)

type BasePage struct {
	// HTMLTitle is the value to use in the page’s <title> tag.
	HTMLTitle string

	// MetaDescription is the html used for rendering the <meta name="Description"> tag.
	MetaDescription safehtml.HTML
}

// servePage is used to execute all templates for a *Server.
func (s *Server) servePage(w http.ResponseWriter, templateName string, page any) {
	buf, err := s.renderPage(templateName, page)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error page", http.StatusInternalServerError)
	}
	if _, err := io.Copy(w, bytes.NewReader(buf)); err != nil {
		log.Println(err)
		http.Error(w, "server error page", http.StatusInternalServerError)
	}
}

// renderPage executes the given templateName with page.
func (s *Server) renderPage(templateName string, page any) ([]byte, error) {
	tmpl, err := s.findTemplate(templateName)
	if err != nil {
		return nil, err
	}
	return executeTemplate(templateName, tmpl, page)
}

// newBasePage returns a base page for the given request and title.
func (s *Server) newBasePage(r *http.Request, title string) BasePage {
	return BasePage{
		HTMLTitle: title,
	}
}

// staticPageHandler handles requests to a template that contains no dynamic
// content.
func (s *Server) staticPageHandler(templateName, title string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.servePage(w, templateName, s.newBasePage(r, title))
	}
}
