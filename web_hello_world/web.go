package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// setup router
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("path", r.URL.Path)
		fmt.Fprintf(w, "pong! on %s\n", r.URL.Path)
	})

	// listen and serve
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
