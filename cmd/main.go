package main

import (
	"e-procurement/internals/initializer"
	"fmt"
	"net/http"
	"os"
)

func main(){
	app, err := initializer.InitializeApp()
	if err != nil {
		panic(err)
	}
	defer app.DB.Close()
	var serverPort string
	if serverPort = os.Getenv("PORT"); serverPort == "" {
		serverPort = "8080" // Default port
	}

	fmt.Printf("Starting server on port %s...\n", serverPort)
	if err := http.ListenAndServe(":"+serverPort, app.Router); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
		panic(err)
	}
}