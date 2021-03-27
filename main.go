package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

//main this is the main application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
