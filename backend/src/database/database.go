package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

const schema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    tasks INTEGER,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

func New() {

	db, err := sql.Open("sqlite", "./app.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(1)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	log.Println("Database is opened.")

	results, err := db.Exec("PRAGMA journal_mode = WAL;")
	fmt.Println(results, err)
	results, err = db.Exec("PRAGMA foreign_keys = ON;")
	fmt.Println(results, err)
	results, err = db.Exec("PRAGMA busy_timeout = 5000;")
	fmt.Println(results, err)

	if _, err := db.Exec(schema); err != nil {
		fmt.Println("database error", err)
		log.Fatal(err)
	}

	fmt.Println("Testing database...")
	// Insert
	res, err := db.Exec("INSERT INTO users (name) VALUES (?)", "Alice")
	if err != nil {

		log.Fatal(err)
	}
	id, _ := res.LastInsertId()

	// Query single row
	var name string
	err = db.QueryRow("SELECT name FROM users WHERE id = ?", id).Scan(&name)
	if err != nil {

		log.Fatal(err)
	}

	// Query multiple rows
	rows, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
