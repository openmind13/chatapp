package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/openmind13/chatapp/utils"
)

// Session struct
type Session struct {
	ID        int
	UUID      string
	UserUUID  string
	CreatedAt time.Time
}

// GetSessionByUUID ...
func GetSessionByUUID(uuid string) (*Session, error) {
	// add code

	return nil, nil
}

// GetUser ...
func (s *Session) GetUser() (*User, error) {
	rows, err := db.Query(`SELECT * FROM users WHERE uuid = $1;`, s.UserUUID)
	if err != nil {
		return nil, errors.New("User from session not found")
	}

	var u User
	err = rows.Scan(&u.ID, &u.UUID, &u.Username, &u.Email, &u.EncryptedPassword, &u.CreatedAt)
	if err != nil {
		fmt.Println("error in GetUser in scanning rows")
	}

	return &u, nil
}

// NewSession ...
func NewSession() *Session {
	return &Session{
		UUID:      utils.GenerateUUID(),
		CreatedAt: time.Now(),
	}
}
