package main

import (
	"errors"
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

// Divide is the divide page handler
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100, 0)
	if err != nil {
		fmt.Fprint(w, fmt.Sprintf("Cannot divide by 0"))
		return
	}
	fmt.Fprint(w, fmt.Sprintf("%f divided by %f is %f",  100.0, 0, f))
}

//addValues adds two integers and return the sum
func addValues(x, y int) int {
	return x + y
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

//main this is the main application
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
