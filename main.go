// ipify-api
//
// This is the main package which starts up and runs our REST API service.
//
// ipify is a simple API service which returns a user's public IP address (it
// supports handling both IPv4 and IPv6 addresses).

package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rdegges/ipify-api/api"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

// main launches our web server which runs indefinitely.
func main() {

	// Setup all routes.  We only service API requests, so this is basic.
	router := httprouter.New()
	router.GET("/", api.GetIP)

	// Setup 404 / 405 handlers.
	router.NotFound = http.HandlerFunc(api.NotFound)
	router.MethodNotAllowed = http.HandlerFunc(api.MethodNotAllowed)

	// Setup middlewares.  For this we're basically adding:
	//	- Support for CORS to make JSONP work.
	handler := cors.Default().Handler(router)

	// Start the server.
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("Starting HTTP server on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))

}
