package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 3)
	_, _ = fmt.Fprint(w, fmt.Sprintf("This is the about page and 2 + 3 is %d", sum))
}

//addValues adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}

//main this is the main application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
