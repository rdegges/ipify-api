// ipify-api
//
// This software implements a basic REST API that provides users with a simple
// way to query their public IP address (IPv4 or IPv6).  This code assumes that
// you are running it on Heroku's platform (https://www.heroku.com/).

package main

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
)

type IPAddress struct {
	IP string `json:"ip"`
}

// getIP returns a user's public facing IP address (IPv4 OR IPv6).
//
// By default, it will return the IP address in plain text, but can also return
// data in both JSON and JSONP if requested to.
func getIP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	ip := net.ParseIP(r.Header["X-Forwarded-For"][len(r.Header["X-Forwarded-For"])-1]).String()

	fmt.Println("DEBUG", r.Form["format"])
	fmt.Println("DEBUG", len(r.Form["format"]))
	if format, ok := r.Form["format"]; ok && len(format) > 0 {
		fmt.Println("DEBUG", format[0])
		switch format[0] {
		case "json":
			jsonStr, _ := json.Marshal(IPAddress{ip})
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, string(jsonStr))
		default:
			w.Header().Set("Content-Type", "text/plain")
			fmt.Fprintf(w, ip)
		}
	}

}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
}

func main() {
	http.HandleFunc("/", getIP)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
