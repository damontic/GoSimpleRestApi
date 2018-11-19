package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	version = "v1.0.0"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s", strings.Replace(r.URL.Path, "/", "", -1))
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Version: %s", version)
}

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	fmt.Fprintf(w, "Hostname: %s", hostname)
}

func main() {

	port := flag.String("port", "8080", "The port in which the microservice is running.")
	flag.Parse()

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/version", versionHandler)
	http.HandleFunc("/hostname", hostnameHandler)
	http.Handle("/metrics", prometheus.Handler())

	log.Printf("Running on port %s\n", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
