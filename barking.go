package bark

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Post struct {
	Id        int
	IdUser    int
	IdComment int
	Text      string
	Likes     int
	Dislikes  int
	Date      string
	Title     string
}

type AllPosts struct {
	Post []Post
}

var post Post
var allPosts AllPosts

//SELECT MAX(id) FROM table;

func getDataBase() *sql.DB {
	db, errSQLOpen := sql.Open("sqlite3", "./public/barkBDD.db")
	if errSQLOpen != nil {
		log.Fatalln(errSQLOpen)
	}

	return db
}

func selectLastId() int {
	db := getDataBase()

	row, errQuery := db.Query("SELECT MAX(id) FROM Post")
	if errQuery != nil {
		log.Fatalln(errQuery)
	}

	var idLastPost int = 0

	for row.Next() {
		err := row.Scan(&idLastPost)
		if err != nil {
			log.Fatal(err)
		}
	}
	row.Close()

	return idLastPost
}

/*
	func selectPost() {
		db, errSQLOpen := sql.Open("sqlite3", "./public/barkBDD.db")
		if errSQLOpen != nil {
			log.Fatalln(errSQLOpen)
		}

		row, errQuery := db.Query("SELECT * FROM Post")
		if errQuery != nil {
			log.Fatalln(errQuery)
		}

		for row.Next() {
			err := row.Scan(&post.Id, &post.IdUser, &post.IdComment, &post.Text, &post.Likes, &post.Dislikes, &post.Date, &post.Title)
			if err != nil {
				log.Fatal(err)
			}
			allPosts.Post = append(allPosts.Post, post)
		}
		row.Close()
	}
*/

func insertPost(insert Post) {

	db := getDataBase()

	statement, errPrepare := db.Prepare("INSERT INTO Post (id, idUser, idComment, text, likes, dislikes, date, title) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if errPrepare != nil {
		log.Fatalln(errPrepare)
	}
	_, errExec := statement.Exec(insert.Id, insert.IdUser, insert.IdComment, insert.Text, insert.Likes, insert.Dislikes, insert.Date, insert.Title)
	if errExec != nil {
		log.Fatalln(errExec)
	}
}

/*
func updatePost(update string) {
	db := getDataBase()

	_, errExec := db.Exec("UPDATE Post SET title = '" + update + "'")
	if errExec != nil {
		log.Fatalln(errExec)
	}
}


func delatePost(delate string) {
	db := getDataBase()

	_, errExec := db.Exec("DELETE FROM Post WHERE title = ''")
	if errExec != nil {
		log.Fatalln(errExec)
	}
}
*/

func getDatePost() string {
	timeNow := time.Now()

	hour := timeNow.Hour()
	minutes := timeNow.Minute()
	day := timeNow.Day()
	month := timeNow.Month()
	year := timeNow.Year()

	var date string = fmt.Sprintf("%d:%d %d/%d/%d", hour, minutes, day, month, year)

	return date
}