package session

import (
	"fmt"
	"net/http"
	"time"

	"github.com/kanciogo/kancio-chat/models"
)

const Length int = 30

var Users = map[string]models.Users{}      // user ID, user
var Sessions = map[string]models.Session{} // session ID, session
var LastCleaned time.Time = time.Now()

func AlreadyLoggedIn(w http.ResponseWriter, req *http.Request) bool {
	ck, err := req.Cookie("session")
	if err != nil {
		return false
	}
	s, ok := Sessions[ck.Value]
	if ok {
		s.LastActivity = time.Now()
		Sessions[ck.Value] = s
	}
	_, ok = Users[s.Username]
	// refresh session
	ck.MaxAge = Length
	http.SetCookie(w, ck)
	return ok
}
func Clean() {
	fmt.Println("BEFORE CLEAN") // for demonstration purposes
	Show()                      // for demonstration purposes
	for k, v := range Sessions {
		if time.Now().Sub(v.LastActivity) > (time.Second * 30) {
			delete(Sessions, k)
		}
	}
	LastCleaned = time.Now()
	fmt.Println("AFTER CLEAN") // for demonstration purposes
	Show()                     // for demonstration purposes
}

// for demonstration purposes
func Show() {
	fmt.Println("********")
	for k, v := range Sessions {
		fmt.Println(k, v.Username)
	}
	fmt.Println("")
}
