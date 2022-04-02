package main

import (
	"fmt"
	"net/http"
)

// const cannot be overriden, and var can
const portNumber = ":8080"

// main is the main application function
func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// start server on port :8080
	// _ = means if there is an error, I don't care
	_ = http.ListenAndServe(portNumber, nil)
}
