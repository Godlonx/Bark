package bark

import (
	"database/sql"
	"fmt"
	"log"
	"math"
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

	defer db.Close()

	row, errQuery := db.Query(request)
	if errQuery != nil {
		log.Fatalln(errQuery)
	}

	for row.Next() {
		var post Post
		err := row.Scan(&post.Id, &post.IdUser, &post.IdComment, &post.Content, &post.Title, &post.Like, &post.Dislike, &post.Date, &post.Tag)
		if err != nil {
			log.Fatal(err)
		}
		currentPosts.Post = append(currentPosts.Post, post)
	}
	row.Close()

	return currentPosts
}

func getPost() {
	db := getDataBase()

	defer db.Close()

	fmt.Println(idPost)

	row, errQuery := db.Query("SELECT * FROM Post WHERE id = '" + idPost + "';")
	if errQuery != nil {
		log.Fatalln(errQuery)
	}

	for row.Next() {
		err := row.Scan(&postClick.Id, &postClick.IdUser, &postClick.IdComment, &postClick.Content, &postClick.Title, &postClick.Like, &postClick.Dislike, &postClick.Date, &postClick.Tag)
		if err != nil {
			log.Fatal(err)
		}
	}
	row.Close()
}

func insertPost(post Post) {

	if post.Title != "" && post.Content != "" {
		db := getDataBase()

		statement, errPrepare := db.Prepare("INSERT INTO Post (id, idUser, idComment, content, title, like, dislike, date, tag) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
		if errPrepare != nil {
			log.Fatalln(errPrepare)
		}
		_, errExec := statement.Exec(post.Id, post.IdUser, post.IdComment, post.Content, post.Title, post.Like, post.Dislike, post.Date, post.Tag)
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
			lastPost = firstPost - 1
			firstPost = lastPost - NUMBER_CURRENT_POSTS + 1
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
		break
	}
}

func getDatePost() string {
	timeNow := time.Now()

	hour := timeNow.Hour()
	minutes := timeNow.Minute()
	day := timeNow.Day()
	monthTmp := timeNow.Month()
	year := timeNow.Year()

	var month string

	switch monthTmp {
	case 1:
		month = "Jan."
		break

	case 2:
		month = "Feb."
		break

	case 3:
		month = "Mar."
		break

	case 4:
		month = "Apr."
		break

	case 5:
		month = "May"
		break

	case 6:
		month = "Jun."
		break

	case 7:
		month = "Jul."
		break

	case 8:
		month = "Aug."
		break

	case 9:
		month = "Sep."
		break

	case 10:
		month = "Oct."
		break

	case 11:
		month = "Nov."
		break

	case 12:
		month = "Dec."
		break
	}

	var date string = fmt.Sprintf("%d %s %d, at %d:%d", day, month, year, hour, minutes)

	return date
}
