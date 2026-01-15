package main

import (
	"fmt"
	"highlevel/connection"
	"highlevel/router"
)

func init() {
	connection.InitDatabase()
}

func main() {
	// Initialize router
	r := router.InitRouter()

	// Start the server
	fmt.Println("Starting HighLevel Service on :8080")
	if err := r.Run(":8080"); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
