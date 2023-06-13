package bark

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

type LoginData struct {
	Username string
	Password string
}

type UserConnected struct {
	Id       int
	Username string
	Password string
	Email    string
	Lvl      int
	Barks    int
	Likes    int
	Dislikes int
}

func Login(user LoginData) (bool, int) {
	db, err := sql.Open("sqlite3", "public/barkBDD.db")
	if err != nil {
		log.Fatalln(err)
	}
	Sql()
	if len(userList.User) > 0 {
		for i := 0; i < len(userList.User); i++ {
			if user.Username == userList.User[i].Username {
				if CheckPasswordHash(user.Password, userList.User[i].Password) {
					return true, userList.User[i].Id
				}

			}
		}
	}
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	return false, 0
}

func SelectUser(id int) UserConnected {
	var user UserConnected

	db, err := sql.Open("sqlite3", "public/barkBDD.db")

	if err != nil {
		log.Fatalln(err)
	}

	row, err := db.Query("select * from user where id=" + strconv.Itoa(id))

	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Lvl, &user.Barks, &user.Likes, &user.Dislikes)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer row.Close()
	return user
}
