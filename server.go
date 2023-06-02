package bark

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var port = ":8080"

// type User struct {
// 	Id       int
// 	Pseudo   string
// 	Password string

// }

type LoginData struct {
	Email    string
	Password string
}

type RegisterData struct {
	Email         string
	Password      string
	Username      string
	Passwordverif string
}

// var user User

func Server() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", ServLogin)
	http.HandleFunc("/home", ServHome)
	http.HandleFunc("/login", ServLogin)
	http.HandleFunc("/register", ServRegister)
	http.HandleFunc("/settings", ServSettings)

	fmt.Println("http://localhost" + port + "/")
	fmt.Println("Server started on port", port)
	http.ListenAndServe(port, nil)
}

func ServHome(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("template/home.html"))
	t.Execute(w, "")
}

func ServLogin(w http.ResponseWriter, r *http.Request) {
	user := Sql()
	t := template.Must(template.ParseFiles("template/login.html"))

	email := r.FormValue("email")
	password := r.FormValue("password")
	data := LoginData{}
	data.Email = email
	data.Password = password
	println(data.Email)
	println(data.Password)
	t.Execute(w, user)
}

func ServRegister(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/register.html"))
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")
	passwordverif := r.FormValue("passwordverif")
	data := RegisterData{}
	data.Email = email
	data.Password = password
	data.Username = username
	data.Passwordverif = passwordverif
	println(data.Username)
	println(data.Email)
	println(data.Password)
	println(data.Passwordverif)
	t.Execute(w, nil)
}

func ServSettings(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/settings.html"))
	t.Execute(w, nil)
}
