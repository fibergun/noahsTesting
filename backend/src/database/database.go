package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"

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
CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task TEXT NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

const logsSchema = `
CREATE TABLE IF NOT EXISTS logs (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

type Database struct {
	*sql.DB
}

type UsersEntry struct {
	id        int
	name      string
	timestamp time.Time
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

	_, err = db.Exec("PRAGMA journal_mode = WAL;")
	if err != nil {
		log.Fatal("setting journal mode: ", err)
	}
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal("setting foreign keys: ", err)
	}
	_, err = db.Exec("PRAGMA busy_timeout = 5000;")
	if err != nil {
		log.Fatal("setting timeout: ", err)
	}

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

	return Database{db}

}

func (db Database) Login(name string) {
	fmt.Println("logging in", name)

	res, err := db.Exec("INSERT INTO users (name) VALUES (?)", name)
	if err != nil {
		log.Fatal("Error inserting user: ", err)
	}
	id, _ := res.LastInsertId()

	//res, err = db.UsersDatabase.Exec("INSERT INTO tasks (task) VALUES (?)", name)
	//if err != nil {
	//	log.Fatal("Error inserting task: ", err)
	//}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users WHERE name = ?", name).Scan(&count)
	if err != nil {
		log.Fatal("Error query-ing count: ", err)
	}
	fmt.Println("User count: ", count)

	var entry UsersEntry
	err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&entry.id, &entry.name, &entry.timestamp)
	fmt.Println(entry, err)
	if err != nil {

		log.Fatal("Error query-ing user: ", err)
	}

}

func (db Database) DeleteAll() {
	res, err := db.Exec("DELETE FROM users")
	if err != nil {
		log.Fatal("Error deleting users: ", err)
	}
	fmt.Println(res)
}
