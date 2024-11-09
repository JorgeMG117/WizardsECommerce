package server

import (
	"crypto/tls"
    "os"
	"fmt"
	"log"
	"net/http"
	"time"
    "github.com/alexedwards/scs/v2"
    "github.com/stripe/stripe-go/v81"

	"github.com/JorgeMG117/WizardsECommerce/routes"
)

func ExecServer() error {
	//Check if CI env variable is set
	// if os.Getenv("DEV") == "true" {
	// 	// Load environment variables from .env file for local development
	// 	if err := godotenv.Load(".env"); err != nil {
	// 		log.Fatalf("Error loading .env file for server: %s %s %s afks", err, os.Getenv("CI"), os.Getenv("DBUSER"))
	// 	}
	// }
    if err := routes.LoadTemplates(); err != nil {
        log.Fatalf("Error loading templates: %v", err)
    }

	s := routes.Server{
		// Db: configs.ConnectDB(),
        SessionManager: scs.New(),
	}

    s.SessionManager.Lifetime = 24 * time.Hour                     // Session lifetime of 24 hours
    s.SessionManager.Cookie.Name = "session_id"                     // Name of the cookie
    s.SessionManager.Cookie.HttpOnly = true                         // Prevent JavaScript access to the cookie
    s.SessionManager.Cookie.Secure = true                           // Ensure the cookie is only sent over HTTPS
    s.SessionManager.Cookie.SameSite = http.SameSiteLaxMode         // Prevent CSRF, but allow navigation from external sites
    s.SessionManager.Cookie.Persist = true                          // Keep the cookie even after the browser is closed




	// defer s.Db.Close()

	// if initializeDB {
	// 	go func() {
	// 		data.InitializeDatabase(s.Db)
	// 		for {
	// 			time.Sleep(time.Hour)
	// 			data.UpdateDatabase(s.Db)
	// 		}
	// 	}()
	// }

	// fmt.Println("Waiting for database to update")
	// time.Sleep(time.Second * 30)

    stripe.Key = os.Getenv("STRIPE_KEY")

	fmt.Println("### CREATING SERVER ###")

	// load tls certificates
	serverTLSCert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
	}
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverTLSCert},
	}

	serv := &http.Server{
		Addr:           ":8080",
		Handler:        s.Router(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		TLSConfig:      tlsConfig,
	}

    fmt.Println("### Server is starting... Listening on https://localhost:8080 ###")

	defer serv.Close()
	log.Fatal(serv.ListenAndServeTLS("", ""))

	return nil
}
