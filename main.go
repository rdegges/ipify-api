// ipify-api
//
// This is the main package which starts up and runs our REST API service.
//
// ipify is a simple API service which returns a user's public IP address (it
// supports handling both IPv4 and IPv6 addresses).

package main

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/rdegges/ipify-api/api"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

// main launches our web server which runs indefinitely.
func main() {

	// Setup all routes.  We only service API requests, so this is basic.
	router := mux.NewRouter()
	router.HandleFunc("/", api.GetIP)

	// Setup middlewares.  For this we're basically adding:
	//	- Support for CORS to make JSONP work.
	n := negroni.New()
	n.Use(cors.Default())
	n.UseHandler(router)

	// Start the server.
	log.Println("Starting HTTP server on port:", os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"), n)
	if err != nil {
		log.Fatal("Error: Couldn't start the HTTP server:", err)
	}

}
