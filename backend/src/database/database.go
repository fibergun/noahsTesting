package database

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "modernc.org/sqlite"
)

const usersSchema = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    group_id INTEGER NOT NULL,
    name TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(group_id, name)
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
    group_id INTEGER NOT NULL,
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
	GroupID   int64     `json:"groupID"`
	Name      string    `json:"username"`
	Timestamp time.Time `json:"timestamp"`
}

type TasksEntry struct {
	ID        int64     `json:"taskID"`
	Task      string    `json:"task"`
	UserID    int64     `json:"userID"`
	GroupID   int64     `json:"groupID"`
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
		log.Println(err)
	}

	_, err = db.Exec("INSERT INTO groups (name) VALUES (?)", "riddersenprincessen")
	if err != nil {
		log.Println(err)
	}

	return Database{db}

}

func (db *Database) GetGroup(group string) (int, error) {
	fmt.Println("Getting group", group)
	var something int

	resp := db.QueryRow("SELECT id FROM groups WHERE name = ?", group)

	err := resp.Scan(&something)
	if err != nil {
		return 0, err
	}

	return something, nil
}

func (db Database) Login(name string, groupID int) (UsersEntry, error) {
	fmt.Println("logging in", name)

	_, err := db.Exec("INSERT INTO users (name, group_id) VALUES (?,?)", name, groupID)
	if err != nil {
		if !strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return UsersEntry{}, err
		}

	}

	return db.GetUser(name, groupID)

}

func (db Database) GetUser(user string, groupID int) (UsersEntry, error) {

	getUser := &UsersEntry{}

	row := db.QueryRow("SELECT id, group_id, name, created_at FROM users WHERE name = ? AND group_id = ?", user, groupID)

	err := row.Scan(&getUser.ID, &getUser.GroupID, &getUser.Name, &getUser.Timestamp)
	if err != nil {
		log.Println("Error getting user: ", err)
		return UsersEntry{}, fmt.Errorf("getting user: %v", err)

	}
	return *getUser, nil

}

func (db Database) GetUserByID(userID int) (UsersEntry, error) {

	getUser := &UsersEntry{}

	row := db.QueryRow("SELECT id, group_id, name, created_at FROM users WHERE id = ?", userID)

	err := row.Scan(&getUser.ID, &getUser.GroupID, &getUser.Name, &getUser.Timestamp)
	if err != nil {
		log.Println("Error getting user: ", err)
		return UsersEntry{}, fmt.Errorf("getting user: %v", err)

	}
	return *getUser, nil

}

func (db Database) InsertTask(task string, user UsersEntry) (TasksEntry, error) {

	fmt.Println("inserting task", task)
	res, err := db.Exec("INSERT INTO tasks (task, user_id, group_id) VALUES (?, ?, ?)", task, user.ID, user.GroupID)
	if err != nil {
		return TasksEntry{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return TasksEntry{}, fmt.Errorf("getting task id: %v", err)
	}

	return db.GetTask(id)

}

func (db Database) GetTask(taskID int64) (TasksEntry, error) {
	getTask := &TasksEntry{}

	row := db.QueryRow("SELECT id, task, user_id, group_id, created_at FROM tasks WHERE id = ?", taskID)

	err := row.Scan(&getTask.ID, &getTask.Task, &getTask.UserID, &getTask.GroupID, &getTask.CreatedAt)
	if err != nil {
		log.Println("Error getting task: ", err)
		return TasksEntry{}, fmt.Errorf("getting task: %v", err)
	}

	return *getTask, nil
}

func (db Database) GetAllTasksByUserID(userID int) ([]TasksEntry, error) {

	rows, err := db.Query("SELECT id, task, user_id, group_id, created_at FROM tasks WHERE user_id = ? ORDER BY id", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	getAllTasks := []TasksEntry{}
	for rows.Next() {
		var task TasksEntry
		err = rows.Scan(&task.ID, &task.Task, &task.UserID, &task.GroupID, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		getAllTasks = append(getAllTasks, task)
	}
	return getAllTasks, nil
}

//func (db Database) getRandomTask(userID int) (TasksEntry, error) {
//
//	user, err := db.GetUserByID(userID)
//	if err != nil {
//		return TasksEntry{}, err
//	}
//
//	//get random from task database
//
//	// make sure they don't have them yet
//
//	//add to log database
//
//	//return which task they have gotten
//
//}

func (db Database) GetRandomTask(groupID int64) (TasksEntry, error) {
	getTask := &TasksEntry{}
	row := db.QueryRow("SELECT id, task, group_id, user_id, created_at FROM tasks WHERE group_id = ? ORDER BY RANDOM() LIMIT 1", groupID)

	err := row.Scan(&getTask.ID, &getTask.Task, &getTask.UserID, &getTask.GroupID, &getTask.CreatedAt)
	if err != nil {
		log.Println("Error getting task: ", err)
		return TasksEntry{}, err
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
//fmt.Println("UserID count: ", count)
//
//var entry UsersEntry
//err = db.QueryRow("SELECT * FROM users WHERE id = ?", id).Scan(&entry.id, &entry.name, &entry.timestamp)
//fmt.Println(entry, err)
//if err != nil {
//
//	log.Fatal("Error query-ing user: ", err)
//}
