package main

import (
	"fmt"
	"log"
	"net/http"
	"flag"
	"strings"
)

const (
	version = "v1.0.0"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", strings.Replace(r.URL.Path, "/", "", -1))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Version: %s", version)
}

func main() {

	port := flag.String("port", "8080", "The port in which the microservice is running.")
	flag.Parse()

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/version", versionHandler)

	log.Println("Running on port %s", *port)
	log.Fatal(http.ListenAndServe(":" + *port, nil))
}
