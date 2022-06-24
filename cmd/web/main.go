package main

import (
	"BBWA/pkg/config"
	"BBWA/pkg/handlers"
	"BBWA/pkg/rander"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

//main is the main application func
func main() {
	var app config.AppConfig

	tc, err := rander.CreateTemplateCash()
	if err != nil {
		log.Fatal("cannot create template cash")
	}

	app.TemplateCash = tc

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
