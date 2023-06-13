package bark

import (
	//"fmt"
	"regexp"
	"unicode"
)

type registerError string

const (
	BadUsername     registerError = "bad username"
	BadPassword                   = "bad password"
	UnequalPassword               = "unequal password"
	BadEmail                      = "bad email"
	Other                         = "Other"
	None                          = "none"
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

func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$") //regex found on stackOverflow
	return emailRegex.MatchString(e)
}
