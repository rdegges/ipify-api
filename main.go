// ipify-api
//
// This is the main package which starts up and runs our REST API service.
//
// ipify is a simple API service which returns a user's public IP address (it
// supports handling both IPv4 and IPv6 addresses).

package main

import (
	"github.com/gorilla/mux"
	"github.com/rdegges/ipify-api/api"
	"log"
	"net/http"
	"os"
)

// main launches our web server which runs indefinitely.
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", api.GetIP)
	http.Handle("/", r)

	log.Println("Starting HTTP server on port:", os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("Error: Couldn't start the HTTP server:", err)
	}
}
