package bark

import (
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var port = ":8080"

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
		err := Login(data)
		if err != nil {
			fmt.Println(err)
			t.Execute(w, err)
		} else {
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
		http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
	}
	t.Execute(w, nil)
}

func ServSettings(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/settings.html"))
	if r.Method == http.MethodPost {
		title := r.FormValue("title")
		switch title {
		case ("password"):
			{
				actualPassword := r.FormValue("actualPassword")
				newPassword := r.FormValue("newPassword")
				validPassword := r.FormValue("validatePassword")
				err := ChangePassword(actualPassword, newPassword, validPassword)
				if err != nil {
					fmt.Println(err)
				}
				break
			}
		case ("username"):
			{
				newUsername := r.FormValue("new")
				err := ChangeUsername(newUsername)
				if err != nil {
					fmt.Println(err)
				}
				break
			}
		case ("email"):
			{
				newEmail := r.FormValue("new")
				err := ChangeEmail(newEmail)
				if err != nil {
					fmt.Println(err)
				}
				break
			}
		case ("delete"):
			{
				DeleteAccount()
				println("AAAAAAAAAAA")
				http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)
			}
		}
	}
	t.Execute(w, user)
}
