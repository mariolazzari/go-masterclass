package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Ciao</h1>"))
}

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>About</h1>"))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Contact</h1>"))
}
