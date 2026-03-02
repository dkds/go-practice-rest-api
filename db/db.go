package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connect to database, " + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)

	createTables()
}

func createTables() {
	createUserTable := `
	CREATE TABLE IF NOT EXISTS user (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		password_salt TEXT NOT NULL
	)
	`
	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("Could not create user table, " + err.Error())
	}

	createEventTable := `
	CREATE TABLE IF NOT EXISTS event (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES user(id)
	)
	`
	_, err = DB.Exec(createEventTable)
	if err != nil {
		panic("Could not create event table, " + err.Error())
	}
}
