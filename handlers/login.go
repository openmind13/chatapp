package handlers

import (
	"fmt"
	"net/http"

	"github.com/openmind13/chatapp/models"
	"github.com/openmind13/chatapp/utils"
	"github.com/openmind13/chatapp/view"
)

// LoginHandler /login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		loginGET(w, r)
	case http.MethodPost:
		loginPOST(w, r)
	default:
		loginGET(w, r)
	}
}

func loginGET(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(userUUIDCookieName)
	if err != nil {
		view.GenerateHTML(w, nil, "login", "header")
		return
	}

	user, err := models.GetUserByUUID(cookie.Value)
	if err != nil {
		// user was not found in the database
		view.GenerateHTML(w, nil, "login", "header")
		return
	}

	view.GenerateHTML(w, user, "login", "header")
}

func loginPOST(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := models.GetUserByEmail(email)
	if err != nil {
		fmt.Println(err)
	}

	if user.EncryptedPassword == utils.CreateHash(password) {
		userUUIDCookie := http.Cookie{
			Name:  userUUIDCookieName,
			Value: user.UUID,
		}
		http.SetCookie(w, &userUUIDCookie)

		usernameCookie := http.Cookie{
			Name:  usernameCookieName,
			Value: user.Username,
		}
		http.SetCookie(w, &usernameCookie)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
