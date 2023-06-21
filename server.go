package bark

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var port = ":8080"

const NUMBER_CURRENT_POSTS = 25

var firstPost = 1
var lastPost = NUMBER_CURRENT_POSTS

var idPost string
var postClick Post
var order string

func Server() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", ServLogin)
	http.HandleFunc("/home", ServHome)
	http.HandleFunc("/topic", ServTopic)
	http.HandleFunc("/login", ServLogin)
	http.HandleFunc("/register", ServRegister)
	http.HandleFunc("/settings", ServSettings)
	http.HandleFunc("/post", ServPost)
	http.HandleFunc("/comment", ServComment)

	fmt.Println("http://localhost" + port + "/")
	fmt.Println("Server started on port", port)
	http.ListenAndServe(port, nil)
}

func ServHome(w http.ResponseWriter, r *http.Request) {

	if user.Username == "" {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
	t := template.Must(template.ParseFiles("template/home.html"))
	var browseDirection string

	if r.Method == http.MethodPost {

		r.ParseForm()
		orderSelect := r.FormValue("select-order")
		if orderSelect == "Earliest" {
			order = "DESC"
		} else if orderSelect == "Latest" {
			order = ""
		}

		browseDirection = r.FormValue("browse-posts")
		browsePosts(browseDirection)

		idPost = r.FormValue("idPost")
		if idPost != "" {
			http.Redirect(w, r, "/topic", http.StatusSeeOther)
		}
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}
	var currentPosts CurrentPosts
	currentPosts = selectTwentyFivePost(firstPost, lastPost, currentPosts, order)

	homeStruct := HomeStruct{currentPosts, user}

	t.Execute(w, homeStruct)
}

func ServTopic(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/topic.html"))
	postId := r.URL.Query().Get("id")
	err, topic := Topic(postId)
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, topic) //post
}

func ServLogin(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/login.html"))
	aliveCookie := false
	dataCookie := getCookie(w, r)
	if dataCookie.Username != "" {
		aliveCookie = true
	}

	if !aliveCookie {
		var data LoginData

		err := Login(data)
		_ = err

		if r.Method == http.MethodPost {
			username := r.FormValue("username")
			password := r.FormValue("password")
			r.ParseForm()
			box := r.Form["remember"]

			data.Username = username
			data.Password = password
			err := Login(data)
			if err != nil {
				fmt.Println(err)
				t.Execute(w, err)
			} else {
				if len(box) == 1 {
					SetCookie(w, r)
				}

				http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
			}
		}
	} else {
		user = dataCookie
		http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
	}

	//
	t.Execute(w, "")
}

func ServRegister(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/register.html"))
	var errRegister registerError = None
	var err error
	isValid := false
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")
		passwordverif := r.FormValue("passwordverif")
		data := RegisterData{}
		data.Email = email
		data.Password = password
		data.Username = username
		data.Passwordverif = passwordverif
		isValid, errRegister = Check(data)

		if isValid {
			err, errRegister = Register(data)
			if err == nil {
				http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
			}

		}
		http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
	}
	t.Execute(w, errRegister)
}

func ServSettings(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("template/settings.html"))
	if r.Method == http.MethodPost {
		disconnect := r.FormValue("disconnect")
		if disconnect == "disconnect" {
			DeleteCookie(w, r)
		}
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
				http.Redirect(w, r, "http://localhost:8080/", http.StatusSeeOther)
			}
		}
	}
	t.Execute(w, user)
}

func ServPost(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/post.html"))

	var tags = GetTag()

	if r.Method == http.MethodPost {

		var post Post
		if tableIsEmpty() {
			post.Id = 1
		} else {
			post.Id = selectLastId() + 1
		}
		post.IdUser = 0
		post.IdComment = 0
		post.Title = r.FormValue("title")
		post.Content = r.FormValue("textarea")
		post.Date = getDatePost()
		post.Likes = 0
		post.Dislikes = 0
		tag := r.FormValue("newTag")
		var idTag int
		if tag != "" {
			tag, idTag = addTag(tag)
			post.Tag = tag
		} else {
			tag = r.FormValue("tag")
			idTag = GetIdTag(tag)
			post.Tag = tag
		}
		insertPost(post, idTag)
		http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
	}
	t.Execute(w, tags)
}

func ServComment(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("template/post.html"))
	var tags = GetTag()
	postId := r.URL.Query().Get("id")
	if r.Method == http.MethodPost {
		var post Post
		if tableIsEmpty() {
			post.Id = 1
		} else {
			post.Id = selectLastId() + 1
		}
		post.IdUser = 0
		post.IdComment = 0
		post.Title = r.FormValue("title")
		post.Content = r.FormValue("textarea")
		post.Date = getDatePost()
		post.Likes = 0
		post.Dislikes = 0
		tag := r.FormValue("newTag")
		var idTag int
		if tag != "" {
			tag, idTag = addTag(tag)
			post.Tag = tag
		} else {
			tag = r.FormValue("tag")
			idTag = GetIdTag(tag)
			post.Tag = tag
		}
		insertComment(post, idTag, postId)
		http.Redirect(w, r, "http://localhost:8080/home", http.StatusSeeOther)
	}

	t.Execute(w, tags)
}
