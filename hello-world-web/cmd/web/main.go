package main

import (
	"fmt"
	"github.com/lukaszczajkowski/golang-udemy/pkg/handlers"
	"net/http"
)

// const cannot be overriden, and var can
const portNumber = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// start server on port :8080
	// _ = means if there is an error, I don't care
	_ = http.ListenAndServe(portNumber, nil)
}
