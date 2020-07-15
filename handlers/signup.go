package handlers

import (
	"fmt"
	"net/http"

	"github.com/openmind13/chatapp/models"
	"github.com/openmind13/chatapp/view"
)

// SignupHandler defines the method
func SignupHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		signupGET(w, r)
	case http.MethodPost:
		signupPOST(w, r)
	default:
		signupGET(w, r)
	}
}

// GET /signup
func signupGET(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(userUUIDCookieName)
	if err != nil {
		view.GenerateHTML(w, nil, "signup", "header")
		return
	}

	user, err := models.GetUserByUUID(cookie.Value)
	if err != nil {
		// user was not found in the database
		view.GenerateHTML(w, nil, "signup", "header")
		return
	}

	view.GenerateHTML(w, user, "signup", "header")
}

// POST /signup
func signupPOST(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")

	err := models.CreateUser(username, email, password)
	if err != nil {
		fmt.Println(err)
	}

	// redirect user to login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
