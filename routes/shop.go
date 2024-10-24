package routes

import (
	"net/http"
	"html/template"
)

func (s *Server) Shop(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("views/shop.html"))
    tmpl.Execute(w, nil)
}

