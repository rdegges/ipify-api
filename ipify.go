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
	host := net.ParseIP(r.Header["X-Forwarded-For"][len(r.Header["X-Forwarded-For"])-1]).String()
	jsonStr, _ := json.MarshalIndent(IPAddress{host}, "", "  ")

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(jsonStr))
}

func main() {
	http.HandleFunc("/", ipify)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}
