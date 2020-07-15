package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/openmind13/chatapp/utils"
)

// User struct
type User struct {
	ID                int
	UUID              string
	Username          string
	Email             string
	EncryptedPassword string
	CreatedAt         time.Time
}

// CreateUser - write user data into database "users"
func CreateUser(username, email, password string) error {
	// check if username already in database
	if checkUserByUsername(username) || checkUserByEmail(email) {
		return errors.New("User already exist")
	}

	_, err := db.Query(`INSERT INTO users(uuid,
										  username, 
										  email,
										  encrypted_password,
										  created_at) 
						VALUES ($1, $2, $3, $4, $5);`,
		utils.GenerateUUID(),
		username,
		email,
		utils.CreateHash(password),
		time.Now(),
	)
	if err != nil {
		return errors.New("Failed to write into DB")
	}

	fmt.Printf("user %v created", username)

	return nil
}

// check by emal if user exists return true, if not exist return false
func checkUserByEmail(email string) bool {
	rows, err := db.Query("SELECT * FROM users WHERE email = $1;", email)
	if err != nil {
		fmt.Println("error in checkUserByEmail")
	}
	defer rows.Close()

	// return true if user is found
	return rows.Next()
}

func checkUserByUsername(username string) bool {
	rows, err := db.Query(`SELECT * FROM users WHERE username = $1;`, username)
	if err != nil {
		fmt.Println("error in checkUserByUsername")
	}
	defer rows.Close()

	return rows.Next()
}

func getUserByUsername(username string) (*User, error) {
	rows, err := db.Query(`SELECT id,
								  uuid,
								  username,
								  email,
								  encrypted_password,
								  created_at
						   FROM users WHERE username = $1;`, username)
	if err != nil {
		fmt.Println("error in getUserByUsername")
		return nil, err
	}

	for rows.Next() {
		user := User{}

		err = rows.Scan(
			&user.ID,
			&user.UUID,
			&user.Username,
			&user.Email,
			&user.EncryptedPassword,
			&user.CreatedAt,
		)
		if err != nil {
			fmt.Println("error in getUserByUsername")
			return nil, err
		}

		return &user, nil
	}

	// user not found in database
	return nil, errors.New("User was not found in the database")
}

// GetUserByEmail ...
func GetUserByEmail(email string) (*User, error) {
	rows, err := db.Query(`SELECT id,
								  uuid,
								  username,
								  email,
								  encrypted_password,
								  created_at
						   FROM users WHERE email = $1;`, email)
	if err != nil {
		fmt.Println("error in getUserByEmail")
		return nil, err
	}

	for rows.Next() {
		user := User{}

		err = rows.Scan(
			&user.ID,
			&user.UUID,
			&user.Username,
			&user.Email,
			&user.EncryptedPassword,
			&user.CreatedAt,
		)
		if err != nil {
			fmt.Println("error in getUserByEmail")
			return nil, err
		}

		return &user, nil
	}

	// user not found in the database
	return nil, errors.New("User was not found in the database")
}

// GetUserByUUID - return user by uuid
func GetUserByUUID(uuid string) (*User, error) {
	rows, err := db.Query(`SELECT id,
								  uuid,
								  username,
								  email,
								  encrypted_password,
								  created_at
							FROM users WHERE uuid = $1;`, uuid)
	if err != nil {
		fmt.Print("error in getUserByUUID\nMake sure that postgresql server is running\n\n")
		return nil, err
	}

	for rows.Next() {
		user := User{}

		err = rows.Scan(
			&user.ID,
			&user.UUID,
			&user.Username,
			&user.Email,
			&user.EncryptedPassword,
			&user.CreatedAt,
		)
		if err != nil {
			fmt.Println("error in getUserByUUID")
			return nil, err
		}

		return &user, nil
	}

	return nil, errors.New("User was not found in the database")
}
