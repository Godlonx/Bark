package bark

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

func getDataBase() *sql.DB {
	db, errSQLOpen := sql.Open("sqlite3", "./public/barkBDD.db")
	if errSQLOpen != nil {
		log.Fatalln(errSQLOpen)
	}

	return db
}

func tableIsEmpty() bool {
	db := getDataBase()
	defer db.Close()

	var rowCount int
	err := db.QueryRow("SELECT COUNT(*) FROM Post").Scan(&rowCount)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if rowCount == 0 {
		return true
	} else {
		return false
	}
}

func selectLastId() int {
	db := getDataBase()
	defer db.Close()

	row, errQuery := db.Query("SELECT MAX(id) FROM Post")
	if errQuery != nil {
		log.Fatalln(errQuery)
		return 0
	}
	defer db.Close()
	defer row.Close()

	var idLastPost int = 0

	for row.Next() {
		err := row.Scan(&idLastPost)
		if err != nil {
			log.Fatal(err)
			return 0
		}
	}

	return idLastPost
}

func selectTwentyFivePost(firstId int, lastId int, currentPosts CurrentPosts) CurrentPosts {
	db := getDataBase()
	defer db.Close()
	var request string = fmt.Sprintf("SELECT * FROM Post WHERE id BETWEEN %d AND %d LIMIT 25", firstId, lastId)

	row, errQuery := db.Query(request)
	if errQuery != nil {
		log.Fatalln(errQuery)
		return CurrentPosts{}
	}
	defer db.Close()
	defer row.Close()

	for row.Next() {
		var post Post
		err := row.Scan(&post.Id, &post.IdUser, &post.IdComment, &post.Content, &post.Title, &post.Likes, &post.Dislikes, &post.Date, &post.Tag)
		if err != nil {
			log.Fatal(err)
			return CurrentPosts{}
		}
		currentPosts.Post = append(currentPosts.Post, post)
	}

	return currentPosts
}

func insertPost(post Post) {

	if post.Title != "" && post.Content != "" {
		db := getDataBase()
		defer db.Close()

		statement, errPrepare := db.Prepare("INSERT INTO Post (id, idUser, idComment, title, content, date, likes, dislikes) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
		if errPrepare != nil {
			log.Fatalln(errPrepare)
			return
		}
		_, errExec := statement.Exec(post.Id, post.IdUser, post.IdComment, post.Title, post.Content, post.Date, post.Likes, post.Dislikes, post.Tag)
		if errExec != nil {
			log.Fatalln(errExec)
			return
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
			firstPost = selectLastId() - NUMBER_CURRENT_POSTS + 1
			lastPost = selectLastId()

		} else {
			firstPost += NUMBER_CURRENT_POSTS
			lastPost += NUMBER_CURRENT_POSTS
		}
		break

	case "last-posts":
		firstPost = selectLastId() - NUMBER_CURRENT_POSTS + 1
		lastPost = selectLastId()
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

func GetTag() []string {
	var Tags []string
	db := getDataBase()

	row, errQuery := db.Query("SELECT * FROM Tag")
	if errQuery != nil {
		log.Fatalln(errQuery)
	}
	defer db.Close()
	defer row.Close()

	for row.Next() {
		var tagName string
		var id int
		err := row.Scan(&id, &tagName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tagName)
		Tags = append(Tags, tagName)
	}
	return Tags
}

func addTag(newtag string) string {
	newtag = strings.ToLower(newtag)
	tags := GetTag()
	for _, tag := range tags {
		if newtag == strings.ToLower(tag) {
			return tag
		}
	}
	db := getDataBase()
	defer db.Close()
	_, err := db.Exec("INSERT INTO tag(name) VALUES(?);", newtag)
	if err != nil {
		log.Fatal(err)
	}
	return newtag
}
