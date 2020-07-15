package handlers

import (
	"net/http"

	"github.com/openmind13/chatapp/models"
	"github.com/openmind13/chatapp/view"
)

// IndexHandler - defines the method for /
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		indexGET(w, r)
	case http.MethodPost:
		indexPOST(w, r)
	default:
		indexGET(w, r)
	}
}

// GET /
func indexGET(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(userUUIDCookieName)
	if err != nil {
		view.GenerateHTML(w, nil, "index", "header")
		return
	}

	user, err := models.GetUserByUUID(cookie.Value)
	if err != nil {
		// user was not found in the database
		view.GenerateHTML(w, nil, "index", "header")
		return
	}

	view.GenerateHTML(w, user, "index", "header")
}

// POST /
func indexPOST(w http.ResponseWriter, r *http.Request) error {
	return nil
}
