package routes

import (
	"net/http"
	"html/template"
)

func (s *Server) Index(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("views/index.html"))
    tmpl.Execute(w, nil)
}
