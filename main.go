package main

import (
	"fmt"
	"os"

	"github.com/leosalgado/keys-sock/client"
	"github.com/leosalgado/keys-sock/server"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Please specity 'server' or 'client'.")
		return
	}

	switch os.Args[1] {
	case "server":
		server.StartServer()
	case "client":
		client.StartClient()
	default:
		fmt.Println("Invalid argument. Please specity 'server' or 'client'.")
	}

}
