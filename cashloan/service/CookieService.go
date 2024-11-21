package service

import (
	"math/rand"
	"net/http"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numberBytes = "0123456789"

func cookieRandomizer() string {
	b := make([]byte, 6)
	for i := range b {
		randomInt := rand.Intn(100)
		if randomInt < 51 {
			b[i] = letterBytes[rand.Intn(len(letterBytes))]
		} else {
			b[i] = numberBytes[rand.Intn(len(numberBytes))]
		}
	}
	return string(b)
}

func SetCookie(w http.ResponseWriter) {
	cookie := http.Cookie{
		Name:     "SessionId",
		Value:    cookieRandomizer(),
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, &cookie)
}

func FindCookie(w http.ResponseWriter, r *http.Request) bool {
	cookie, err := r.Cookie("sessionId")
	if err != nil {
		return false
	} else {
		_ = cookie.Value
		return true
	}

}
