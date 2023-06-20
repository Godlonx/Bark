package bark

import (
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

var test User

func Login(userLogin LoginData) error {

	rows := getData("SELECT * From user where username = '" + userLogin.Username + "'")
	for rows.Next() {

		err := rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Lvl, &user.Barks, &user.Likes, &user.Dislikes)
		return err
	}
	defer rows.Close()
	return errors.New("No existing account")

}
