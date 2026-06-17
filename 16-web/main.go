package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)

	fmt.Println("Server started on port 8081")
	if err := http.ListenAndServe(":8081", mux); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

}
