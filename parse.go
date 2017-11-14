package main 

import (
	"net/http"
	"fmt"
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
	Output(w, chat)
} // chatHandler

func Output(w http.ResponseWriter, chat string){
	fmt.Fprintf(w, " %s", chat);
} // Output