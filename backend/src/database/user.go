package database

import (
	"fmt"
	"log"
	"strings"
)

func (db Database) Login(name string, groupID int) (UsersEntry, error) {
	fmt.Println("logging in", name)

	_, err := db.Exec("INSERT INTO users (name, group_id) VALUES (?,?)", name, groupID)
	if err != nil {
		if !strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return UsersEntry{}, err
		}

	}

	return db.GetUser(name)

}

func (db Database) GetUser(user string) (UsersEntry, error) {

	getUser := &UsersEntry{}

	row := db.QueryRow("SELECT id, group_id, name, created_at FROM users WHERE name = ?", user)

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
