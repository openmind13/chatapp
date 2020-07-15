package handlers

import (
	"net/http"

	"github.com/openmind13/chatapp/models"
	"github.com/openmind13/chatapp/view"
)

// ChatHandler view the chat page /
func ChatHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(userUUIDCookieName)
	if err != nil {
		view.GenerateHTML(w, nil, "chat", "header")
		return
	}

	user, err := models.GetUserByUUID(cookie.Value)
	if err != nil {
		// user was not found in the database
		view.GenerateHTML(w, nil, "chat", "header")
		return
	}

	view.GenerateHTML(w, user, "chat", "header")
}
