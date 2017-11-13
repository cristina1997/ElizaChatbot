package main
//Code: https://getbootstrap.com/docs/4.0/components/forms/
import (
	"net/http"
	//"text/template"
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
}

func main() {
	http.HandleFunc("/", indexHTML)
	//http.HandleFunc("/chat/", chatHandler)

	http.ListenAndServe(":9999", nil)
}
