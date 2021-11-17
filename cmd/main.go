package main

import (
	"fmt"

	"github.com/RamiroCuenca/users-crud-test/routes"
	"github.com/RamiroCuenca/users-crud-test/server"
)

func main() {
	// Get router
	r := routes.GetRouter()

	// Get server
	sv := server.GetServer(r)

	// Run server
	fmt.Println("Server running through port :8080...")
	sv.Run()
}
