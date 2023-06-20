package bark

import (
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var test User

func Login(userLogin LoginData) error {
	fmt.Println(userLogin.Username)
	rows := getData("SELECT * From user where username = '" + userLogin.Username + "'")
	for rows.Next() {
		fmt.Println("aaaaaaaaaaa")
		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Lvl, &user.Barks, &user.Likes, &user.Dislikes)
		if CheckPasswordHash(userLogin.Password, user.Password) {
			return err
		}
	}
	defer rows.Close()
	return errors.New("No existing account")

}
