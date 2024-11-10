package routes

import (
	"net/http"

	"github.com/JorgeMG117/WizardsECommerce/models"
)

func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
        s.RenderTemplate(w, "register.html", nil)
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
