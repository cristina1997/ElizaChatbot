package main

import (
	"net/http"
	
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/chat/", chatHandler)

	http.ListenAndServe(":9999", nil)
} // main
