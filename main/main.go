package main

import (
	"fmt"
	"html/template"
	"database/sql"
	_"github.com/mattn/go-sqlite3"
	"net/http"
	"log"
)

var port = ":8080"

type User struct {
	Id	int
	Pseudo string
	Password string
}


func main() {
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/home", Home)

	fmt.Println("(http://localhost"+port+"/home"+") - Server started on port", port)
	http.ListenAndServe(port, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	user := Sql()
	t := template.Must(template.ParseFiles("template/home.html"))
	t.Execute(w,user)
}

func Sql() User{


	user := User{Id :0,Pseudo:"",Password:""}


	db, err := sql.Open("sqlite3", "bark.db")
	
	if err != nil {
		log.Fatalln(err)
	}


	rows, err := db.Query("select * from user")
	
	if err != nil {
		log.Fatal(err)
	}
	
	for rows.Next() {
		err := rows.Scan(&user.Id,&user.Pseudo,&user.Password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(user)

	}

	defer rows.Close()

	return user
}