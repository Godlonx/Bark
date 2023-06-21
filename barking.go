package bark

import (
	"database/sql"
	"fmt"
	"log"
	"math"
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

func selectTwentyFivePost(firstId int, lastId int, currentPosts CurrentPosts,order string) CurrentPosts {
	db := getDataBase()
	defer db.Close()

	var request string = fmt.Sprintf("SELECT * FROM Post WHERE id BETWEEN %d AND %d ORDER BY date "+order+" LIMIT 25", firstId, lastId)

	defer db.Close()

	row, errQuery := db.Query(request)
	if errQuery != nil {
		log.Fatalln(errQuery)
		return CurrentPosts{}
	}
	defer row.Close()

	for row.Next() {
		var post Post
		err := row.Scan(&post.Id, &post.IdUser, &post.IdComment, &post.Content, &post.Title, &post.Likes, &post.Dislikes, &post.Date)
		if err != nil {
			log.Fatal(err)
			return CurrentPosts{}
		}
		currentPosts.Post = append(currentPosts.Post, post)
	}
	println(len(currentPosts.Post))
	for i := 0; i < len(currentPosts.Post); i++ {
		row, err := db.Query("SELECT tag.name FROM Post JOIN tagRef on Post.id = tagRef.idPost JOIN tag on tagRef.idTag = tag.id WHERE Post.id = ?;", currentPosts.Post[i].Id)
		if err != nil {
			log.Fatal(err)
		}
		for row.Next() {
			err := row.Scan(&currentPosts.Post[i].Tag)
			println(currentPosts.Post[i].Id, currentPosts.Post[i].Tag)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return currentPosts
}

func insertPost(post Post, idTag int) {

	if post.Title != "" && post.Content != "" {
		db := getDataBase()
		defer db.Close()

		query, errPrepare := db.Prepare("INSERT INTO Post (id, idUser, idComment, title, content, date, like, dislike) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
		if errPrepare != nil {
			log.Fatalln(errPrepare)
			return
		}
		_, errExec := query.Exec(post.Id, user.Id, post.IdComment, post.Title, post.Content, post.Date, post.Likes, post.Dislikes)
		if errExec != nil {
			log.Fatalln(errExec)
			return
		}

		_, err := db.Exec("INSERT INTO tagRef(idTag, idPost) VALUES(?, ?)", idTag, post.Id)
		if err != nil {
			log.Fatal(err)
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
		Tags = append(Tags, tagName)
	}
	return Tags
}

func addTag(newtag string) (string, int) {
	newtag = strings.ToLower(newtag)
	tags := GetTag()
	lastId := 0
	for id, tag := range tags {
		if newtag == strings.ToLower(tag) {
			return tag, id
		}
		lastId = id
	}
	db := getDataBase()
	defer db.Close()
	_, err := db.Exec("INSERT INTO tag(name) VALUES(?);", newtag)
	if err != nil {
		log.Fatal(err)
	}
	return newtag, lastId + 1
}

func GetIdTag(searchedTag string) int {
	tags := GetTag()
	for id, tag := range tags {
		if strings.ToLower(searchedTag) == strings.ToLower(tag) {
			return id
		}
	}
	return 0
}
