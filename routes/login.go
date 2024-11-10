package routes

import (
	"fmt"
	"net/http"
    "strconv"

	"github.com/JorgeMG117/WizardsECommerce/models"
)

func (s *Server) LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
        s.RenderTemplate(w, "login.html", nil)
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

    fmt.Println("User login: ", user.ID)
    userIdString := strconv.FormatUint(uint64(user.ID), 10)


    s.SessionManager.Put(r.Context(), "user_id", userIdString)


    http.Redirect(w, r, "/", http.StatusSeeOther)
}

