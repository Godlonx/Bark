package main

import (
	"fmt"
	"html/template"	
	"net/http"
	"bark"

)

var port = ":8080"




func main() {
	bark.Register()
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/home", Home)

	fmt.Println("(http://localhost"+port+"/home"+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("template/home.html"))
	t.Execute(w,"")
}

