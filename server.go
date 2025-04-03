package main

import(
	"fmt"
	"net/http"
)

func TemplateIndex(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Test Tempalte HTML Golang.")
}

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Test / Golang")
	})
	http.HandleFunc("/home", TemplateIndex)
	
	fmt.Println("Server is running in http://127.0.0.1:2001")
	http.ListenAndServe(":2001", nil)
}