package bark

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
)

func ChangeUsername(newUsername string) error {
	db, err := sql.Open("sqlite3", "./public/barkBDD.db")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer db.Close()
	fmt.Println(newUsername, user.Username)
	_, err = db.Exec("UPDATE User SET username = ? WHERE username = ?;", newUsername, user.Username)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	user.Username = newUsername
	return nil
}

func ChangeEmail(newMail string) error {
	if !isEmailValid(newMail) {
		return errors.New("Invalid email syntax")
	}
	val := ""
	rows := getData("Select email From User")
	for rows.Next() {
		err := rows.Scan(&val)
		if err != nil {
			println(err)
		}
		if val == newMail {
			return errors.New("Already existing email")
		}
	}
	rows.Close()
	user.Email = newMail
	sendData("UPDATE user SET email = '" + newMail + "' WHERE id='" + strconv.Itoa(user.Id) + "'")
	return nil
}

func ChangePassword(actualPassword string, newPassword string, validPassword string) error {
	if newPassword != validPassword {
		return errors.New("validation password different")
	}
	if !verifyPassword(newPassword) {
		return errors.New("Invalid password syntax")
	}
	hashPassword, err := HashPassword(newPassword)
	if err != nil {
		return err
	}
	sendData("Update user Set password = " + hashPassword + " Where id=" + strconv.Itoa(user.Id) + "")
	return nil
}

func DeleteAccount() {
	sendData("Delete From user Where username='" + user.Username + "'")
}
