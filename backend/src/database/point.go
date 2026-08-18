package database

import (
	"fmt"
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

	return 2*completed - total, nil
}
