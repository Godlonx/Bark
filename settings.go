package bark

import (
	"database/sql"
	"errors"
	"log"
	"strconv"
)

func ChangeUsername(newUsername string) error {
	/*val := ""
	rows := getData("Select username From User")
	for rows.Next() {
		err := rows.Scan(&val)
		print("tttttttttttttttttt")
		if err != nil {
			println(err)
		}
		if val == newUsername {
			return errors.New("Already existing username")
		}
	}
	rows.Close()
	print(user.Id)*/
	db, err := sql.Open("sqlite3", "file:public/barkBDD.db?cache=shared")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	query, err := db.Prepare("UPDATE User SET username = '" + newUsername + "' WHERE id = '" + strconv.Itoa(user.Id) + "';")
	if err != nil {
		log.Fatalln(err)
	}
	_, err = query.Exec()
	if err != nil {
		log.Fatalln(err)
	}
	println("UPDATE User SET username = '" + newUsername + "' WHERE id = '" + strconv.Itoa(user.Id) + "';")
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
	user.Email = newMail
	sendData("UPDATE user SET email = " + newMail + " WHERE id=" + strconv.Itoa(user.Id) + "")
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
