package routes

import (
	"net/http"
	"html/template"
)

func (s *Server) Contact(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("views/contact.html"))
    tmpl.Execute(w, nil)
}

