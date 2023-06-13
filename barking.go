package bark

import (
	"database/sql"
	"log"
)

type Post struct {
	Title     string
	NbReplies int
	Timestamp string
	Tag       string
	Content   string
	Like      int
	Dislike   int
}

type AllPosts struct {
	Post []Post
}

var post Post
var allPosts AllPosts

func selectPost() {
	db, errSQLOpen := sql.Open("sqlite3", "./bark.db")
	if errSQLOpen != nil {
		log.Fatalln(errSQLOpen)
	}

	row, errQuery := db.Query("SELECT * FROM Post")
	if errQuery != nil {
		log.Fatalln(errQuery)
	}

	for row.Next() {
		err := row.Scan(&post.Title, &post.NbReplies, &post.Timestamp, &post.Tag, &post.Content, &post.Like, &post.Dislike)
		if err != nil {
			log.Fatal(err)
		}
		allPosts.Post = append(allPosts.Post, post)
	}
	row.Close()
}

func updatePost(update string) {
	db, errSQLOpen := sql.Open("sqlite3", "./bark.db")
	if errSQLOpen != nil {
		log.Fatalln(errSQLOpen)
	}

	_, errExec := db.Exec("UPDATE Post SET title = '" + update + "'")
	if errExec != nil {
		log.Fatalln(errExec)
	}
}

func insertPost(insert Post) {
	db, errSQLOpen := sql.Open("sqlite3", "./bark.db")
	if errSQLOpen != nil {
		log.Fatalln(errSQLOpen)
	}

	statement, errPrepare := db.Prepare("INSERT INTO Post (title, number_replies, timestamp, tag, content, number_like, number_dislike) VALUES(?, ?, ?, ?, ?, ?, ?)")
	if errPrepare != nil {
		log.Fatalln(errPrepare)
	}
	_, errExec2 := statement.Exec(insert.Title, insert.NbReplies, insert.Timestamp, insert.Tag, insert.Content, insert.Like, insert.Dislike)
	if errExec2 != nil {
		log.Fatalln(errExec2)
	}
}

func delatePost(delate string) {
	db, errSQLOpen := sql.Open("sqlite3", "./bark.db")
	if errSQLOpen != nil {
		log.Fatalln(errSQLOpen)
	}

	_, errExec := db.Exec("DELETE FROM Post WHERE title = ''")
	if errExec != nil {
		log.Fatalln(errExec)
	}
}
