package bark

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var port = ":8080"

var userConnected UserConnected

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
	t.Execute(w, nil)
}

func ServLogin(w http.ResponseWriter, r *http.Request) {
	//user := Sql()
	t := template.Must(template.ParseFiles("template/login.html"))
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		data := LoginData{}
		data.Username = username
		data.Password = password
		authorize, idUser := Login(data)
		println(authorize)
		println(idUser)
		if authorize {
			userConnected = SelectUser(idUser)
			http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
		}
	}
	t.Execute(w, "")
}

func ServRegister(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/register.html"))
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")
		passwordverif := r.FormValue("passwordverif")
		data := RegisterData{}
		data.Email = email
		data.Password = password
		data.Username = username
		data.Passwordverif = passwordverif
		isValid, err := Check(data)
		println(err)
		if isValid {
			Register(data)
		}
	}
	t.Execute(w, nil)
}

func ServSettings(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/settings.html"))
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		if title == "password" {
			actualPassword := r.FormValue("actualPassword")
			newPassword := r.FormValue("newPassword")
			validPassword := r.FormValue("validatePassword")
			println(actualPassword, newPassword, validPassword)
		} else {
			new := r.FormValue("new")
			println(new)
		}
	}
	t.Execute(w, nil)
}
