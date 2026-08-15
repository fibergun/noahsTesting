package database

import "fmt"

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
