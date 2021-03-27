package main

import (
	"fmt"
	"github.com/aldolushkja/gohello-web/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

//main this is the main application
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}
