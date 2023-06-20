package bark

import (
	//"fmt"
	"database/sql"
	"log"
	"regexp"
	"unicode"
)

func Check(registerData RegisterData) (bool, registerError) {
	if !verifyPassword(registerData.Password) {
		return false, BadPassword
	}
	if registerData.Password != registerData.Passwordverif {
		return false, UnequalPassword
	}
	if !isEmailValid(registerData.Email) {
		println("bad")
		return false, BadEmail
	}
	return true, None
}

func verifyPassword(s string) bool {
	letters := 0
	number := false
	upper := false
	special := false
	sevenOrMore := false
	for _, c := range s {
		letters++
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		case unicode.IsLetter(c):
		default:
			return false
		}
	}
	sevenOrMore = letters >= 7
	if number && upper && special && sevenOrMore {
		return true
	}
	return false
}

// regex found on stackOverflow
func isEmailValid(email string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegex.MatchString(email)
}

func isEmailAlreadyUsed(email string) bool {
	existingEmail := ""
	rows := getData("Select email From User")
	for rows.Next() {
		err := rows.Scan(&existingEmail)
		if err != nil {
			println(err)
		}
		if existingEmail == email {
			return true
		}
	}
	return false
}

func isUsernameAlreadyUsed(username string) bool {
	existingUsername := ""
	rows := getData("Select username From User")
	for rows.Next() {
		err := rows.Scan(&existingUsername)
		if err != nil {
			println(err)
		}
		if existingUsername == username {
			return true
		}
	}
	return false
}

func sendData(request string) {
	db, err := sql.Open("sqlite3", "file:public/barkBDD.db?cache=shared")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	query, err := db.Prepare(request)
	if err != nil {
		log.Fatalln(err)
	}
	query.Exec()
}

func getData(request string) *sql.Rows {
	db, err := sql.Open("sqlite3", "public/barkBDD.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()
	rows, err := db.Query(request)
	if err != nil {
		log.Fatal(err)
	}
	return rows
}
