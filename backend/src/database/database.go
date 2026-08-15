package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

//go:embed schemas/*.sql
var schemaFS embed.FS

var schemaFiles = []string{
	"schemas/groups.sql",
	"schemas/tasks.sql",
	"schemas/users.sql",
	"schemas/logs.sql",
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

	for _, file := range schemaFiles {
		b, err := schemaFS.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		_, err = db.Exec(string(b))
		if err != nil {
			log.Fatal(err, string(b))
		}
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

func (db Database) DeleteAll() {
	res, err := db.Exec("DELETE FROM users")
	if err != nil {
		log.Fatal("Error deleting users: ", err)
	}
	fmt.Println(res)
}
