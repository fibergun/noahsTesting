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

const groupsSchema = `
CREATE TABLE IF NOT EXISTS groups (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL UNIQUE
);
`

const tasksSchema = `
CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task TEXT NOT NULL UNIQUE,
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

type TasksEntry struct {
	ID        int64     `json:"taskID"`
	Task      string    `json:"task"`
	UserID    int64     `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
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

	if _, err := db.Exec(groupsSchema); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec(tasksSchema); err != nil {

		log.Fatal(err)
	}

	if _, err := db.Exec(logsSchema); err != nil {

		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO groups (name) VALUES (?)", "brakcie")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("INSERT INTO groups (name) VALUES (?)", "riddersenprincessen")
	if err != nil {
		log.Fatal(err)
	}

	// Insert

	return Database{db}

}

func (db Database) Login(name string) (UsersEntry, error) {
	fmt.Println("logging in", name)

	_, _ = db.Exec("INSERT INTO users (name) VALUES (?)", name)

	return db.GetUser(name)

}

func (db Database) GetUser(user string) (UsersEntry, error) {

	getUser := &UsersEntry{}

	row := db.QueryRow("SELECT id, name, created_at FROM users WHERE name = ?", user)

	err := row.Scan(&getUser.ID, &getUser.Name, &getUser.Timestamp)
	if err != nil {
		log.Println("Error getting user: ", err)
		return UsersEntry{}, fmt.Errorf("getting user: %v", err)

	}
	return *getUser, nil

}

func (db Database) InsertTask(task string, UserID int64) (TasksEntry, error) {

	fmt.Println("inserting task", task)
	res, err := db.Exec("INSERT INTO tasks (task, user_id) VALUES (?, ?)", task, UserID)
	if err != nil {
		return TasksEntry{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return TasksEntry{}, fmt.Errorf("getting task id: %v", err)
	}

	return TasksEntry{
		ID:     id,
		Task:   task,
		UserID: UserID,
	}, nil

}

func (db Database) GetTask(taskID int64) (TasksEntry, error) {
	getTask := &TasksEntry{}

	row := db.QueryRow("SELECT id, task, user_id, created_at FROM tasks WHERE id = ?", taskID)

	err := row.Scan(&getTask.ID, &getTask.Task, &getTask.UserID, &getTask.CreatedAt)
	if err != nil {
		log.Println("Error getting task: ", err)
		return TasksEntry{}, fmt.Errorf("getting task: %v", err)
	}

	return *getTask, nil
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
