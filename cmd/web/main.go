package main

import (
	"BBWA/pkg/handlers"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

//main is the main application func
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
