package main

import (
	"io"
	"net/http"
	"time"
    "log"
    "fmt"
    "crypto/tls"

	"github.com/alexedwards/scs/v2"
)

var sessionManager *scs.SessionManager

func main() {
	// Initialize a new session manager and configure the session lifetime.
	sessionManager = scs.New()
    sessionManager.Lifetime = 24 * time.Hour                     // Session lifetime of 24 hours
    sessionManager.Cookie.Name = "session_id"                     // Name of the cookie
    sessionManager.Cookie.HttpOnly = true                         // Prevent JavaScript access to the cookie
    sessionManager.Cookie.Secure = true                           // Ensure the cookie is only sent over HTTPS
    sessionManager.Cookie.SameSite = http.SameSiteLaxMode         // Prevent CSRF, but allow navigation from external sites
    sessionManager.Cookie.Persist = true                          // Keep the cookie even after the browser is closed

	mux := http.NewServeMux()
	mux.HandleFunc("/put", putHandler)
	mux.HandleFunc("/get", getHandler)

	// Wrap your handlers with the LoadAndSave() middleware.
	//http.ListenAndServe(":4000", sessionManager.LoadAndSave(mux))

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
		Handler:        sessionManager.LoadAndSave(mux),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
		TLSConfig:      tlsConfig,
	}

    fmt.Println("### Server is starting... Listening on https://localhost:8080 ###")

	defer serv.Close()
	log.Fatal(serv.ListenAndServeTLS("", ""))

}

func putHandler(w http.ResponseWriter, r *http.Request) {
	// Store a new key and value in the session data.
	sessionManager.Put(r.Context(), "message", "Hello from a session!")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	// Use the GetString helper to retrieve the string value associated with a
	// key. The zero value is returned if the key does not exist.
	msg := sessionManager.GetString(r.Context(), "message")
	io.WriteString(w, msg)
}
