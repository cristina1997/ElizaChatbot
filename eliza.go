package main

import (
	"net/http"
	"text/template"
	//"math/rand"
	//"time"
)

func checkError(e error) {
    if e != nil {
        panic(e)
    }
}//checkError

type tempData struct {
	Eliza string
}

func indexHTML(w http.ResponseWriter, r * http.Request){
	http.ServeFile(w, r, "src/index.html")

	tmpl, err := template.ParseFiles("src/index.html")
	tmpl.Execute(w, tempData{Eliza: "Hi"})

	checkError(err)
}

func main() {

	http.HandleFunc("/", indexHTML)
	//http.HandleFunc("/chat/", chatHandler)

	http.ListenAndServe(":9999", nil)
}