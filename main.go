package main

import (
	"fmt"
	"github.com/avishkar-activex/chms-auth/services"
)

func main() {
	// start auth service
	fmt.Println("starting auth service")

	// connect with db

	// start the service
	ws := services.NewWebService()

	ws.Start()

	fmt.Println("started auth service")
}
