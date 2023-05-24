package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var port = ":8080"

type User struct {
	Id       int
	Pseudo   string
	Password string
}

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

var user User

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register", Register)

	fmt.Println("(http://localhost"+port+"/login"+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
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

func Register(w http.ResponseWriter, r *http.Request) {
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

func Sql() User {

	db, err := sql.Open("sqlite3", "public/barkData.db")

	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("select * from user")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Pseudo, &user.Password)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer rows.Close()

	return user
}
