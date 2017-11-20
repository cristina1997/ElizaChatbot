package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/chat/", chatHandler)

	http.ListenAndServe(":9999", nil)
} // main

func chatHandler(w http.ResponseWriter, r * http.Request) {
	userInput := r.URL.Query().Get("input")
	elizaResponse := eliza.Ask(userInput)

	for {
        statement := userInput()
        fmt.Println("Eliza: " + eliza.ReplyTo(statement))

        if eliza.IsQuitStatement(statement) {
            break
        }
    }
} // chatHandler
