package bark

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"time"
)

type Post struct {
	Id        int
	IdUser    int
	IdComment int
	Title     string
	Content   string
	Date      string
	Likes     int
	Dislikes  int
}

type CurrentPosts struct {
	Post []Post
}

func getDataBase() *sql.DB {
	db, errSQLOpen := sql.Open("sqlite3", "./public/barkBDD.db")
	if errSQLOpen != nil {
		log.Fatalln(errSQLOpen)
	}

	return db
}

func tableIsEmpty() bool {
	db := getDataBase()

	var rowCount int
	err := db.QueryRow("SELECT COUNT(*) FROM Post").Scan(&rowCount)
	if err != nil {
		log.Fatal(err)
	}

	if rowCount == 0 {
		return true
	} else {
		return false
	}
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

func selectTwentyFivePost(firstId int, lastId int, currentPosts CurrentPosts) CurrentPosts {
	db := getDataBase()

	var request string = fmt.Sprintf("SELECT * FROM Post WHERE id BETWEEN %d AND %d LIMIT 25", firstId, lastId)

	row, errQuery := db.Query(request)
	if errQuery != nil {
		log.Fatalln(errQuery)
	}

	for row.Next() {
		var post Post
		err := row.Scan(&post.Id, &post.IdUser, &post.IdComment, &post.Title, &post.Content, &post.Date, &post.Likes, &post.Dislikes)
		if err != nil {
			log.Fatal(err)
		}
		currentPosts.Post = append(currentPosts.Post, post)
	}
	row.Close()

	return currentPosts
}

func insertPost(post Post) {

	if post.Title != "" && post.Content != "" {
		db := getDataBase()

		statement, errPrepare := db.Prepare("INSERT INTO Post (id, idUser, idComment, title, content, date, likes, dislikes) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
		if errPrepare != nil {
			log.Fatalln(errPrepare)
		}
		_, errExec := statement.Exec(post.Id, post.IdUser, post.IdComment, post.Title, post.Content, post.Date, post.Likes, post.Dislikes)
		if errExec != nil {
			log.Fatalln(errExec)
		}
	}
}

func browsePosts(browseDirection string) {

	switch browseDirection {
	case "first-posts":
		firstPost = 1
		lastPost = NUMBER_CURRENT_POSTS
		break

	case "prev-posts":

		whatIdFirstPostBelow := firstPost - NUMBER_CURRENT_POSTS

		if whatIdFirstPostBelow <= 0 {
			firstPost = 1
			lastPost = NUMBER_CURRENT_POSTS

		} else {
			firstPost -= NUMBER_CURRENT_POSTS
			lastPost -= NUMBER_CURRENT_POSTS
		}
		break

	case "next-posts":

		whatIdLastPostAbove := lastPost + NUMBER_CURRENT_POSTS

		if whatIdLastPostAbove >= selectLastId() {

			howManyCurrentPostsAlreadyRead := math.Round(float64(selectLastId()/NUMBER_CURRENT_POSTS)) * NUMBER_CURRENT_POSTS

			firstPost = int(howManyCurrentPostsAlreadyRead) + 1
			lastPost = selectLastId()

		} else {
			firstPost += NUMBER_CURRENT_POSTS
			lastPost += NUMBER_CURRENT_POSTS
		}
		break

	case "last-posts":
		howManyCurrentPostsAlreadyRead := math.Round(float64(selectLastId()/NUMBER_CURRENT_POSTS)) * NUMBER_CURRENT_POSTS

		firstPost = int(howManyCurrentPostsAlreadyRead) + 1
		lastPost = selectLastId()
		/*
			firstPost = selectLastId() - NUMBER_CURRENT_POSTS + 1
			lastPost = selectLastId()
		*/
		break
	}
}

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
