package bark

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type UserList struct {
	User []User
}

type User struct {
	Id       int
	Username   string
	Password string
	Email    string
}



type RegisterData struct {
	Email         string
	Password      string
	Username      string
	Passwordverif string
}

var user User
var userList UserList

func Register() {

	db, err := sql.Open("sqlite3", "public/barkData.db")

	if err != nil {
		log.Fatalln(err)
	}

	Username := "hi"
	password := "test"
	email := "mathis@ynov.com"
	hashPass, _ := HashPassword(password)
	Sql()

	for i := 0; i < len(userList.User); i++ {

		if userList.User[i].Username == Username {
			fmt.Println("error")
		}

	}

	insert := "INSERT into user (Username,password,email) VALUES ('" + Username + "','" + hashPass + "','" + email + "')"

	_, err = db.Exec(insert)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}

func Sql() User {

	db, err := sql.Open("sqlite3", "public/barkData.db")

	if err != nil {
		log.Fatalln(err)
	}

	rows, err := db.Query("select * from user")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		userList.User = append(userList.User, user)
	}

	defer rows.Close()
	return userList.User[0]
}
