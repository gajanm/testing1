package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/on", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The sprinkler is on")
	})

	http.HandleFunc("/off", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "The sprinkler is off")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
