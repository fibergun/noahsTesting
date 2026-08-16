package database

import (
	"fmt"
	"log"
)

func (db Database) GetPoints(userID int) (int64, error) {
	var total int64
	err := db.QueryRow(`SELECT COUNT(*) FROM logs WHERE user_id = ?`, userID).Scan(&total)
	if err != nil {
		return 0, fmt.Errorf("error getting total tasks: %w", err)
	}

	var completed int64

	err = db.QueryRow(`SELECT COUNT(*) FROM logs WHERE user_id = ? AND completed = true`, userID).Scan(&completed)
	if err != nil {
		return 0, fmt.Errorf("error getting completed tasks: %w", err)
	}

	log.Printf("Total tasks: %d", total)
	log.Printf("Total completed: %d", completed)

	return 2*completed - total, nil
}
