package routes

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/JorgeMG117/WizardsECommerce/models"
)

func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		s.getRegister(w, r)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	username, password := r.FormValue("username"), r.FormValue("password")
	if username == "" || password == "" {
		http.Error(w, "Username and password required", http.StatusBadRequest)
		return
	}

	user := models.CheckUser(username, password)
	if user == nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

}

func (s *Server) getRegister(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("views", "register.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Failed to load login form", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Failed to render login form", http.StatusInternalServerError)
	}
}
