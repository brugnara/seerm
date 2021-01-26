package main

import (
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", index)
	http.Handle("/c/", customer{})
	http.Handle("/favicon.ico", http.NotFoundHandler())
}

func listen(port string) {
	if port == "" {
		port = "8080"
	}
	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
