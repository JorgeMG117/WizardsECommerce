package main

import (
	"fmt"
	"os"
    "log"

	"github.com/JorgeMG117/WizardsECommerce/server"
    "github.com/joho/godotenv"
)

func main() {
    env := os.Getenv("GO_ENV")
    if env != "prod" {
        err := godotenv.Load(".env")
        if err != nil {
            log.Fatal("Error loading .env file")
        }
    }
	//Launch server
	if err := server.ExecServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
