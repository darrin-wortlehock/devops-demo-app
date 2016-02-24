package main

import (
	"fmt"
	"net/http"
	"os"
)

var Build string = "dev"
var Addr string = ":8484"

func handler(w http.ResponseWriter, r *http.Request) {
	h, _ := os.Hostname()
	fmt.Fprintf(w, "Hi there, I'm served from %s!", h)
}

func main() {
	fmt.Printf("Starting devops-demo-app (build: %s) on %s\n", Build, Addr)
	http.HandleFunc("/", handler)
	http.ListenAndServe(Addr, nil)
}