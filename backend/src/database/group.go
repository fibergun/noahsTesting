package database

import "fmt"

func (db *Database) GetGroup(group string) (int64, error) {
	fmt.Println("Getting group", group)
	var something int64

	resp := db.QueryRow("SELECT id FROM groups WHERE name = ?", group)

	err := resp.Scan(&something)
	if err != nil {
		return 0, err
	}

	return something, nil
}
