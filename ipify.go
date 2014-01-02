package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

type IPAddress struct {
	IP string `json:"ip"`
}

func ipify(w http.ResponseWriter, r *http.Request) {
	host, _, err := net.SplitHostPort(r.Header["X-Forwarded-For"][len(r.Header["X-Forwarded-For"])-1])
	if err != nil {
		log.Fatal("SplitHostPort:", err)
	}

	w.Header().Set("Content-Type", "application/json")

	jsonStr, _ := json.MarshalIndent(IPAddress{host}, "", "  ")
	fmt.Fprintf(w, string(jsonStr))
}

func main() {
	http.HandleFunc("/", ipify)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
