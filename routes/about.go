package routes

import (
	"net/http"
	"html/template"
)

func (s *Server) About(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("views/about.html"))
    tmpl.Execute(w, nil)
}

