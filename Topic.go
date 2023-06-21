package bark

import (
	"database/sql"
	"log"
)

func Topic(idPost string) (error, TopicStruct) {
	db, err := sql.Open("sqlite3", "./public/barkBDD.db")
	var topic TopicStruct
	if err != nil {
		return err, TopicStruct{}
	}
	defer db.Close()

	row, err := db.Query("SELECT * FROM Post WHERE id = ?", idPost)
	if err != nil {
		return err, TopicStruct{}
	}
	defer row.Close()
	for row.Next() {
		var post Post
		err := row.Scan(&post.Id, &post.IdUser, &post.IdComment, &post.Content, &post.Title, &post.Likes, &post.Dislikes, &post.Date)
		if err != nil {
			return err, TopicStruct{}
		}
		topic.Post = post
	}

	rowTag, err := db.Query("SELECT tag.name FROM Post JOIN tagRef on Post.id = tagRef.idPost JOIN tag on tagRef.idTag = tag.id WHERE Post.id = ?;", topic.Post.Id)
	if err != nil {
		log.Fatal(err)
	}
	for rowTag.Next() {
		err := rowTag.Scan(&topic.Post.Tag)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer rowTag.Close()

	row2, err := db.Query("SELECT * FROM Post WHERE idComment = ?", idPost)
	if err != nil {
		return err, TopicStruct{}
	}
	defer row2.Close()
	var posts []Post
	for row2.Next() {
		var post Post
		err := row2.Scan(&post.Id, &post.IdUser, &post.IdComment, &post.Content, &post.Title, &post.Likes, &post.Dislikes, &post.Date)
		if err != nil {
			return err, TopicStruct{}
		}
		posts = append(posts, post)
	}

	row3, err := db.Query("SELECT u.* FROM User u JOIN Post p on p.idUser = u.id WHERE p.id = ?", idPost)
	if err != nil {
		return err, TopicStruct{}
	}
	defer row3.Close()
	for row3.Next() {
		var user2 User
		err := row3.Scan(&user2.Id, &user2.Username, &user2.Password, &user2.Email, &user2.Lvl, &user2.Barks, user2.Likes, user2.Dislikes)
		if err != nil {
			return err, TopicStruct{}
		}
		topic.User2 = user2
	}

	row4, err := db.Query("SELECT isLike From Likes l JOIN Post p ON p.id = l.idPost WHERE p.id = ?", idPost)
	defer row4.Close()
	for row4.Next() {
		var isLike int
		err := row4.Scan(&isLike)
		if err != nil {
			return err, TopicStruct{}
		}
		if isLike == 1 {
			topic.Post.Likes++
		} else if isLike == -1 {
			topic.Post.Dislikes++
		}
	}

	row5, err := db.Query("SELECT isLike From Likes l JOIN Post p ON p.id = l.idPost JOIN user u on u.id = l.idUser  WHERE p.id = ? AND u.id = ?", idPost, user.Id)
	defer row5.Close()
	for row5.Next() {
		err := row5.Scan(&topic.isLike)
		if err != nil {
			return err, TopicStruct{}
		}
	}

	topic.Comments = posts
	topic.UserConnected = user
	println(topic.User2.Username)
	return nil, topic
}

func insertComment(post Post, idTag int, postId string) {

	if post.Title != "" && post.Content != "" {
		db := getDataBase()
		defer db.Close()

		query, errPrepare := db.Prepare("INSERT INTO Post (id, idUser, idComment, title, content, date, like, dislike) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
		if errPrepare != nil {
			log.Fatalln(errPrepare)
			return
		}
		_, errExec := query.Exec(post.Id, post.IdUser, postId, post.Title, post.Content, post.Date, post.Likes, post.Dislikes)
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
