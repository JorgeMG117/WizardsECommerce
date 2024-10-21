package routes

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/JorgeMG117/WizardsECommerce/models"
)

func (s *Server) UsersPage(w http.ResponseWriter, r *http.Request) {

	tmpl := template.Must(template.ParseFiles("views/users.html"))
	tmpl.Execute(w, nil)
}

func (s *Server) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
    // Retrieve user information from the session
    userID := s.SessionManager.GetInt(r.Context(), "user_id")

    fmt.Println(userID)

	s.mutex.Lock()

	users, _ := models.GetUsers()

	s.mutex.Unlock()

	tmpl := template.Must(template.New("users").Parse(`
		{{range .}}
			<div class="col-md-4 mb-4">
				<div class="card">
					<div class="card-body">
						<h5 class="card-title">{{.Username}}</h5>
						<p class="card-text">{{.Email}}</p>
						<p class="card-text">{{.Role}}</p>
					</div>
				</div>
			</div>
		{{end}}
	`))

	w.Header().Set("Content-Type", "text/html")
	tmpl.Execute(w, users)
}
