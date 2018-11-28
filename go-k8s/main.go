package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request: %#v", r)
		fmt.Fprintf(w, "Hello %s", r.URL.Query().Get("name"))
	})
	log.Fatal(http.ListenAndServe(":18080", nil))
}
