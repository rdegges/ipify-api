// ipify-api
//
// This is the main package which starts up and runs our REST API service.
//
// ipify is a simple API service which returns a user's public IP address (it
// supports handling both IPv4 and IPv6 addresses).

package main

import (
	"github.com/gorilla/mux"
	"github.com/rdegges/ipify-api/controllers"
	"net/http"
	"os"
)

// main launches our web server which runs indefinitely.
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.GetIP)
	http.Handle("/", r)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
