package database

import "fmt"

func (db Database) GetPoints(userID int) (int64, error) {

	total, err := db.Exec(`SELECT COUNT(*) FROM logs WHERE user_id = ?`, userID)
	if err != nil {
		return 0, fmt.Errorf("error getting total tasks: %w", err)
	}
	totalTasks, err := total.RowsAffected()

	completed, err := db.Exec(`SELECT COUNT(*) FROM logs WHERE user_id = ? AND completed = true`, userID)
	if err != nil {
		return 0, fmt.Errorf("error getting completed tasks: %w", err)
	}

	totalCompleted, err := completed.RowsAffected()

	return 2*totalCompleted - totalTasks, nil
}
