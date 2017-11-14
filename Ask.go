package main 

import (
	"net/http"
	"fmt"
)

func Ask(w http.ResponseWriter,chat string){
	fmt.Fprintf(w, " %s", chat);
} // Output