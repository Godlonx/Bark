package bark

import (
	"errors"

	_ "github.com/mattn/go-sqlite3"
)

func Register(data RegisterData) (error, registerError) {
	username := data.Username
	password := data.Password
	email := data.Email
	hashPass, _ := HashPassword(password)
	if isEmailAlreadyUsed(email) {
		return errors.New("Email already used"),AlredyUsedEmail
	}
	if isUsernameAlreadyUsed(username) {
		return errors.New("Username already used"),AlredyUsedUsername
	}
	insert := "INSERT into user (username,password,email,lvl,barks,likes,dislikes) VALUES ('" + username + "','" + hashPass + "','" + email + "','0', '0', '0', '0')"
	sendData(insert)
	return nil,None
}
