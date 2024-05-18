package main 

import (
	"fmt"
	"log"
	"net/http"
	"html"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter,r *http.Request){
		fmt.Fprintf(w, "Heelo, %q", html.EscapeString(r.URL.Path))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":3001", nil))
}