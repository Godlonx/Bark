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
	BadUsername        registerError = "bad username"
	BadPassword                      = "bad password"
	UnequalPassword                  = "unequal password"
	BadEmail                         = "bad email"
	AlredyUsedUsername               = "name already used"
	AlredyUsedEmail                  = "email already used"
	Other                            = "Other"
	None                             = "none"
)

type Post struct {
	Id        int
	IdUser    int
	IdComment int
	Title     string
	Content   string
	Date      string
	Likes     int
	Dislikes  int
	Tag       string
}

type CurrentPosts struct {
	Post []Post
}

type HomeStruct struct {
	Post          CurrentPosts
	UserConnected User
	Tags          []string
}

type TopicStruct struct {
	Post          Post
	Comments      []Post
	UserConnected User
	User2         User
}

var user User
