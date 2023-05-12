package bark

import (
	"fmt"
	"database/sql"
	_"github.com/mattn/go-sqlite3"
	"log"
	"golang.org/x/crypto/bcrypt"
)

type UserList struct{
	User []User
}

type User struct {
	Id	int
	Pseudo string
	Password string
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
	hashPass,_ := HashPassword(password)
	Sql()
	


	for i := 0; i < len(userList.User); i++ {

		fmt.Println(CheckPasswordHash(password,userList.User[i].Password))
		
		// if (userList.User[i].Pseudo== Pseudo){
		// 	fmt.Println("error")
		// }
		// if (userList.User[i].Password== hashPass){
		// 	fmt.Println("error password")
		// }else{
		// 	fmt.Println("good password")
		// }
	}

	

	insert :="INSERT into user (pseudo,password) VALUES ('"+Pseudo+"','"+hashPass+"')"

	_, err = db.Exec(insert)


	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
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
		err := rows.Scan(&user.Id,&user.Pseudo,&user.Password)
		if err != nil {
			log.Fatal(err)
		}
		userList.User = append(userList.User,user)
	}
	
	defer rows.Close()

}