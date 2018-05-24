package main

import (
	"fmt"
	"net/http"
	"html"
	"log"
)

const (
	version = "v1.0.0"
)

func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func healthHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(200)
}

func versionHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Version: %s", version)
}

func main(){
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/version", versionHandler)

	log.Fatal(http.ListenAndServe(":8080",nil))
}
