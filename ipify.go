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

func jsonip(w http.ResponseWriter, r *http.Request) {
	host := net.ParseIP(r.Header["X-Forwarded-For"][len(r.Header["X-Forwarded-For"])-1]).String()
	jsonStr, _ := json.Marshal(IPAddress{host})

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonStr))
}

func textip(w http.ResponseWriter, r *http.Request) {
	host := net.ParseIP(r.Header["X-Forwarded-For"][len(r.Header["X-Forwarded-For"])-1]).String()

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, host)
}

func main() {
	http.HandleFunc("/json", jsonip)
	http.HandleFunc("/text", textip)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
