package bark

type LoginData struct {
	Username string
	Password string
}

type UserList struct {
	User []User
}

type User struct {
	Id       int
	Username string
	Password string
	Email    string
	Lvl      int
	Barks    int
	Likes    int
	Dislikes int
}

type RegisterData struct {
	Email         string
	Password      string
	Username      string
	Passwordverif string
}

type registerError string

const (
	BadUsername     registerError = "bad username"
	BadPassword                   = "bad password"
	UnequalPassword               = "unequal password"
	BadEmail                      = "bad email"
	Other                         = "Other"
	None                          = "none"
)

var user User
