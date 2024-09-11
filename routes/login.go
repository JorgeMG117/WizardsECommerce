package routes

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/JorgeMG117/WizardsECommerce/models"
	"github.com/JorgeMG117/WizardsECommerce/utils"
)

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		s.serveLoginForm(w, r)
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

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (s *Server) serveLoginForm(w http.ResponseWriter, r *http.Request) {
	tmplPath := filepath.Join("views", "login.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Failed to load login form", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Failed to render login form", http.StatusInternalServerError)
	}
}
