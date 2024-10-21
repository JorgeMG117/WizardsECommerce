package middleware

import (
	"fmt"
	"net/http"
    "strconv"

	"github.com/alexedwards/scs/v2"
)

func AuthenticationMiddleware(sessionManager *scs.SessionManager, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // userID := sessionManager.GetInt(r.Context(), "user_id")
        userIDString := sessionManager.GetString(r.Context(), "user_id")
        var userID int
        var err error
        if userIDString == "" { 
            userID = 0
        } else {
            userID, err = strconv.Atoi(userIDString)
            if err != nil {
                fmt.Println("Error:", err)
                return
            }
        }
        fmt.Println("In middleware: ", userID)
        if userID == 0 {
            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }
        next.ServeHTTP(w, r)
	})
}
