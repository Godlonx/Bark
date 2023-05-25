package bark

import (
	"fmt"
	"database/sql"
	_"github.com/mattn/go-sqlite3"
	"log"
)

func Login() []User {

	db, err := sql.Open("sqlite3", "public/barkData.db")
	
	if err != nil {
		log.Fatalln(err)
	}

	Pseudo := "hi"
	password := "test"
	//hashPass,_ := HashPassword(password)
	Sql()
	

	if len(userList.User)>0 {
		for i := 0; i < len(userList.User); i++ {


			if (CheckPasswordHash(password,userList.User[i].Password) && Pseudo == userList.User[i].Pseudo){
				fmt.Println(userList.User[i])
				
			}
		}
	}
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	return userList.User
}