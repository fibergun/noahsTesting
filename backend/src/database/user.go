package database

import (
	"fmt"
	"log"
	"strings"
)

func (db Database) Login(name string, groupID int64) (UsersEntry, error) {
	fmt.Println("logging in", name)

	_, err := db.Exec("INSERT INTO users (name, group_id) VALUES (?,?)", name, groupID)
	if err != nil {
		if !strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return UsersEntry{}, err
		}

	}

	return db.GetUser(name, groupID)

}

func (db Database) GetUser(user string, groupID int64) (UsersEntry, error) {

	getUser := &UsersEntry{}

	row := db.QueryRow("SELECT id, group_id, name, created_at FROM users WHERE (name, group_id) = (?,?)", user, groupID)

	err := row.Scan(&getUser.ID, &getUser.GroupID, &getUser.Name, &getUser.CreatedAt)
	if err != nil {
		log.Println("Error getting user: ", err)
		return UsersEntry{}, fmt.Errorf("getting user: %v", err)

	}
	return *getUser, nil

}

func (db Database) GetUserByID(userID int) (UsersEntry, error) {

	getUser := &UsersEntry{}

	row := db.QueryRow("SELECT id, group_id, name, created_at FROM users WHERE id = ?", userID)

	err := row.Scan(&getUser.ID, &getUser.GroupID, &getUser.Name, &getUser.CreatedAt)
	if err != nil {
		log.Println("Error getting user: ", err)
		return UsersEntry{}, fmt.Errorf("getting user: %v", err)

	}
	return *getUser, nil

}
