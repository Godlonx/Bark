package bark

import (
	"fmt"
	"database/sql"
	_"github.com/mattn/go-sqlite3"
	"log"
)

type UserList struct{
	User []User
}

type User struct {
	Id	int
	Pseudo string
	Password string
	Email string
}
var user User
var userList UserList

func Register() {

	db, err := sql.Open("sqlite3", "public/barkData.db")
	
	if err != nil {
		log.Fatalln(err)
	}

	Pseudo := "hi"
	password := "test"
	email := "mathis@ynov.com"
	hashPass,_ := HashPassword(password)
	Sql()
	


	for i := 0; i < len(userList.User); i++ {

		if (userList.User[i].Pseudo== Pseudo){
			fmt.Println("error")
		}
		
	}

	

	insert :="INSERT into user (pseudo,password,email) VALUES ('"+Pseudo+"','"+hashPass+"','"+email+"')"

	_, err = db.Exec(insert)


	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}



func Sql(){

	db, err := sql.Open("sqlite3", "public/barkData.db")
	
	if err != nil {
		log.Fatalln(err)
	}


	rows, err := db.Query("select * from user")
	
	if err != nil {
		log.Fatal(err)
	}
	
	for rows.Next() {
		err := rows.Scan(&user.Id,&user.Pseudo,&user.Password,&user.Email)
		if err != nil {
			log.Fatal(err)
		}
		userList.User = append(userList.User,user)
	}
	
	defer rows.Close()

}

