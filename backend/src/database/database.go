package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

const usersSchema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

const tasksSchema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task TEXT NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

const logsSchema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

type Database struct {
	UsersDatabase *sql.DB
	tasksDatabase *sql.DB
	logsDatabase  *sql.DB
}

type UsersEntry struct {
	id   int
	name string
}

func New() Database {

	db, err := sql.Open("sqlite", "./app.db")
	if err != nil {
		panic(err)
	}

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

	if _, err := db.Exec(usersSchema); err != nil {

		log.Fatal(err)
	}

	if _, err := db.Exec(tasksSchema); err != nil {

		log.Fatal(err)
	}

	if _, err := db.Exec(logsSchema); err != nil {

		log.Fatal(err)
	}

	// Insert

	return Database{UsersDatabase: db}

}

func (db Database) Login(name string) {
	res, err := db.UsersDatabase.Exec("INSERT INTO users (name, user) VALUES (?, ?)", name, 1)
	if err != nil {
		log.Fatal("Error inserting user: ", err)
	}
	id, _ := res.LastInsertId()

	var count int
	err = db.UsersDatabase.QueryRow("SELECT COUNT(*) FROM users WHERE name = ?", name).Scan(&count)
	if err != nil {
		log.Fatal("Error query-ing count: ", err)
	}
	fmt.Println("User count: ", count)

	// Query single row
	var entry UsersEntry
	err = db.UsersDatabase.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&entry.id, &entry.name)
	fmt.Println(entry, err)
	if err != nil {

		log.Fatal("Error query-ing user: ", err)
	}

}
