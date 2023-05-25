package main

import (
	"fmt"
	"html/template"	
	"net/http"
	"bark"

)

var port = ":8080"




func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/home", Home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register", Register)

	fmt.Println("(http://localhost"+port+"/home"+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("template/home.html"))
	t.Execute(w,"")
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := bark.Login()
	t := template.Must(template.ParseFiles("template/home.html"))
	t.Execute(w,user)
}

func Register(w http.ResponseWriter, r *http.Request) {
	bark.Register()
	t := template.Must(template.ParseFiles("template/home.html"))
	t.Execute(w,"")
}


