package main

import (
	"net/http"
	"fmt"
	"./eliza"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/chat/", chatHandler)

	http.ListenAndServe(":9999", nil)
} // main

func chatHandler(w http.ResponseWriter, r * http.Request) {
	input := r.URL.Query().Get("input")

	elizaResponse := ""			

	if eliza.IsStartStatement(input){
		elizaResponse = eliza.ElizaHi()
	} else if eliza.IsEndStatement(input) {
		elizaResponse = eliza.ElizaBye()
	} else {
		elizaResponse = eliza.Ask(input)
	}	

	fmt.Fprintf(w, elizaResponse);

} // chatHandler
