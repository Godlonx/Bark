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

type FormData struct {
	Email    string
	Password string
}

var user User

func main() {

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/home", Home)

	fmt.Println("(http://localhost"+port+"/home"+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	user := Sql()
	t := template.Must(template.ParseFiles("template/login.html"))

	email := r.FormValue("email")
	password := r.FormValue("password")
	data := FormData{}
	data.Email = email
	data.Password = password
	println(data.Email)
	println(data.Password)
	t.Execute(w, user)
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
