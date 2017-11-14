package main 

import (
	"net/http"
)

func checkError(e error) {
	// If there IS an error 
    if e != nil {
        panic(e) // crash program
    }
} // checkError

func chatHandler(w http.ResponseWriter, r * http.Request) {
	err := r.ParseForm() 
	checkError(err)

	chat := r.FormValue("input") // parse form name="input" from user
	Ask(w, chat)
} // chatHandler

