package main

import (
	"bark"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var port = ":8080"

var userConnected bark.UserConnected

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", Login)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register", Register)

	fmt.Println("http://localhost" + port + "/")
	fmt.Println("Server started on port", port)
	http.ListenAndServe(port, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("template/home.html"))
	println(userConnected.Username)
	t.Execute(w, userConnected)
}

func Login(w http.ResponseWriter, r *http.Request) {
	//user := bark.Sql()
	t := template.Must(template.ParseFiles("template/login.html"))

	username := r.FormValue("username")
	password := r.FormValue("password")
	data := bark.LoginData{}
	data.Username = username
	data.Password = password
	authorize,idUser := bark.Login(data)
	println(authorize)
	println(idUser)
	if (authorize) {
		userConnected=bark.SelectUser(idUser)
		http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
	}
	t.Execute(w, "")
}

func Register(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/register.html"))
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	passwordverif := r.FormValue("passwordverif")
	data := bark.RegisterData{}
	data.Email = email
	data.Password = password
	data.Username = username
	data.Passwordverif = passwordverif

	isValid,err := bark.Check(data)
	println(err)
	if (isValid) {
		
	}
	t.Execute(w, nil)
}
