package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/JorgeMG117/WizardsECommerce/utils"
)

func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header missing", http.StatusUnauthorized)
			return
		}

		fmt.Println(authHeader)

		tokenString := strings.Split(authHeader, " ")[1]
		_, err := utils.VerifyToken(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		//r.Header.Set("user_id", claims["user_id"])
		next.ServeHTTP(w, r)
	})
}

// // AuthenticationMiddleware checks if the user has a valid JWT token
// func AuthenticationMiddleware(next http.HandlerFunc) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		tokenString := c.GetHeader("Authorization")
// 		if tokenString == "" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authentication token"})
// 			c.Abort()
// 			return
// 		}

// 		// The token should be prefixed with "Bearer "
// 		tokenParts := strings.Split(tokenString, " ")
// 		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
// 			c.Abort()
// 			return
// 		}

// 		tokenString = tokenParts[1]

// 		claims, err := utils.VerifyToken(tokenString)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authentication token"})
// 			c.Abort()
// 			return
// 		}

// 		c.Set("user_id", claims["user_id"])
// 		c.Next()
// 	}
// }
