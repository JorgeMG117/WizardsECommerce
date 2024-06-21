package main

import (
	"fmt"
	"os"

	"github.com/JorgeMG117/WizardsECommerce/server"
)

func main() {
	//Launch server
	if err := server.ExecServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
