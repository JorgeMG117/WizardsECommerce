package routes

import (
    "net/http"
)

func (s *Server) logoutHandler(w http.ResponseWriter, r *http.Request) {
    s.SessionManager.Remove(r.Context(), "user_id")
    http.Redirect(w, r, "/login", http.StatusSeeOther)
}
