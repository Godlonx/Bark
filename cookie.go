package bark

import (
    "encoding/base64"
    "errors"
    "net/http"
	"encoding/gob"
	"bytes"
	"log"
	"strings"
)

var (
    ErrValueTooLong = errors.New("cookie value too long")
    ErrInvalidValue = errors.New("invalid cookie value")
)

func SetCookie(w http.ResponseWriter, r *http.Request){
	gob.Register(&User{})
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:     "connectedUser",
		Value:    buf.String(),
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
	}

	err = Write(w, cookie)

	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
}

func getCookie(w http.ResponseWriter, r *http.Request) User{
	data := User{}
	gobEncodedValue, err :=Read(r, "connectedUser")
    if err != nil {
        switch {
        case errors.Is(err, http.ErrNoCookie):
            data.Username = ""
            return data
        case errors.Is(err, ErrInvalidValue):
            http.Error(w, "invalid cookie", http.StatusBadRequest)
        default:
            log.Println(err)
            http.Error(w, "server error", http.StatusInternalServerError)
        }
        return data
    }

    reader := strings.NewReader(gobEncodedValue)

    if err := gob.NewDecoder(reader).Decode(&data); err != nil {
        log.Println(err)
        //http.Error(w, "server error", http.StatusInternalServerError)
        return data
    }
	return data
}

func Write(w http.ResponseWriter, cookie http.Cookie) error {
    cookie.Value = base64.URLEncoding.EncodeToString([]byte(cookie.Value))

    if len(cookie.String()) > 4096 {
        return ErrValueTooLong
    }

    http.SetCookie(w, &cookie)

    return nil
}

func Read(r *http.Request, name string) (string, error) {
    cookie, err := r.Cookie(name)
    if err != nil {
        return "", err
    }

    value, err := base64.URLEncoding.DecodeString(cookie.Value)
    if err != nil {
        return "", ErrInvalidValue
    }

    return string(value), nil
}

func DeleteCookie(w http.ResponseWriter, r *http.Request){
    gob.Register(&User{})
	var buf bytes.Buffer
	err := gob.NewEncoder(&buf).Encode(&user)
	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:     "connectedUser",
		Value:    buf.String(),
		Path:     "/",
		MaxAge:   -1,
	}

	err = Write(w, cookie)

	if err != nil {
		log.Println(err)
		http.Error(w, "server error", http.StatusInternalServerError)
		return
	}

    http.Redirect(w, r, "http://localhost:8080/login", http.StatusSeeOther)
}