package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Yo, it me"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Print("Starting server on :420")
	err := http.ListenAndServe(":420", mux)
	log.Fatal(err)
}
