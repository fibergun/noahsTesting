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
    task TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

const logsSchema = `
CREATE TABLE IF NOT EXISTS logs (
    log_id INTEGER PRIMARY KEY,
    task_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
`

type Database struct {
	*sql.DB
}

type UsersEntry struct {
	ID        int64     `json:"userID"`
	Name      string    `json:"username"`
	Timestamp time.Time `json:"timestamp"`
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

func (db Database) Login(name string) (UsersEntry, error) {
	fmt.Println("logging in", name)

	res, err := db.Exec("INSERT INTO users (name) VALUES (?)", name)
	if err != nil {
		return db.GetUser(name)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error getting user id: ", err)
	}

	return UsersEntry{
		ID:   id,
		Name: name,
	}, nil

}

func (db Database) GetUser(user string) (UsersEntry, error) {
	fmt.Println("getting user", user)

	getUser := &UsersEntry{}

	row := db.QueryRow("SELECT id, name, created_at FROM users WHERE name = ?", user)

	err := row.Scan(&getUser.ID, &getUser.Name, &getUser.Timestamp)
	fmt.Println("Rows: ", row)
	if err != nil {
		log.Println("Error getting user: ", err)
		if err == sql.ErrNoRows {
			return UsersEntry{}, err
		}
	}
	fmt.Println("Get user: ", getUser)
	return *getUser, nil

}

func (db Database) DeleteAll() {
	res, err := db.Exec("DELETE FROM users")
	if err != nil {
		log.Fatal("Error deleting users: ", err)
	}
	fmt.Println(res)
}

//res, err = db.UsersDatabase.Exec("INSERT INTO tasks (task) VALUES (?)", name)
//if err != nil {
//	log.Fatal("Error inserting task: ", err)
//}

//var count int
//err = db.QueryRow("SELECT COUNT(*) FROM users WHERE name = ?", name).Scan(&count)
//if err != nil {
//	log.Fatal("Error query-ing count: ", err)
//}
//fmt.Println("User count: ", count)
//
//var entry UsersEntry
//err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&entry.id, &entry.name, &entry.timestamp)
//fmt.Println(entry, err)
//if err != nil {
//
//	log.Fatal("Error query-ing user: ", err)
//}
