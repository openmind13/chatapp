package models

import (
	"database/sql"
	"log"

	// importing driver for postgresql
	_ "github.com/lib/pq"
)

// database which used by app
var db *sql.DB

var dbLoginStr = "postgres://openmind:228pirog228@localhost/chatdb"

func init() {
	var err error

	db, err = sql.Open("postgres", dbLoginStr)
	if err != nil {
		log.Fatal("Can't open db", err)
	}
}
