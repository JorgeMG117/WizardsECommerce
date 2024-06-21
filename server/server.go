package server

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"

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

	s := routes.Server{
		// Db: configs.ConnectDB(),
	}
	// defer s.Db.Close()

	// if intializeDB {
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

	defer serv.Close()
	log.Fatal(serv.ListenAndServeTLS("", ""))

	return nil
}
