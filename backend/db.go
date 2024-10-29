package main

import (
	"database/sql"

	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS messages (id INTEGER PRIMARY KEY AUTOINCREMENT, message TEXT , username TEXT)")
	if err != nil {
		log.Fatal(err)
	}
}

func OpenDB() *sql.DB {

	_, err := os.Stat("./db.db")

	if errors.Is(err, os.ErrNotExist) {
		log.Println("Database is empty, creating new one...")
		os.Create("./db.db")
	}

	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		log.Fatal(err)
	}

	InitDB(db)
	return db
}

func AddMessage(db *sql.DB, message string, username string) {
	stmt, err := db.Prepare("INSERT INTO messages(message, username) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(message, username)
	if err != nil {
		log.Fatal(err)
	}
}

func GetMessages(db *sql.DB) []string {
	rows, err := db.Query("SELECT message FROM messages")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var messages []string
	for rows.Next() {
		var message string
		if err := rows.Scan(&message); err != nil {
			log.Fatal(err)
		}
		messages = append(messages, message)
	}
	return messages
}

func CloseDB(db *sql.DB) {
	db.Close()
}
