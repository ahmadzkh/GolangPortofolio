/**
 * File main.go
 *
 * This is the main entry point for the web application.
 * It sets up the HTTP server, routes, and handlers.
 *
 * @ahmadzkh
 */
package main

import (
	"ProjectLepkom/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/assets/",
		http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	mux.HandleFunc("/", handler.HomeHandler)

	log.Println("Starting Port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
