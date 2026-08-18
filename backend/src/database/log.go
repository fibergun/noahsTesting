package database

import (
	"database/sql"
	"fmt"
)

func (db Database) addTaskToLogs(taskID int64, userID int64) (LogsEntry, error) {
	checkTask, ok, err := db.checkTaskExistsInLogs(taskID, userID, false)
	if ok {
		return checkTask, fmt.Errorf("user already has this task")
	}
	if err != nil {
		return LogsEntry{}, err
	}

	_, err = db.Exec("INSERT INTO logs (task_id, user_id) VALUES (?, ?) ", taskID, userID)
	if err != nil {
		return LogsEntry{}, fmt.Errorf("failed to add task to logs: %v", err)
	}

	task, ok, err := db.checkTaskExistsInLogs(taskID, userID, true)
	if !ok {
		return LogsEntry{}, fmt.Errorf("failed to retrieve task from logs: %v", err)
	}
	return task, nil
}

func (db Database) checkTaskExistsInLogs(taskID int64, userID int64, errOnEmpty bool) (LogsEntry, bool, error) {
	getLog := LogsEntry{}

	row := db.QueryRow("SELECT * FROM logs WHERE (task_id, user_id) = (?,?)", taskID, userID)

	err := row.Scan(&getLog.ID, &getLog.TaskID, &getLog.UserID, &getLog.Completed, &getLog.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			if errOnEmpty {
				return LogsEntry{}, false, fmt.Errorf("task not found: %v", err)
			}
			return LogsEntry{}, false, nil

		}
		return LogsEntry{}, false, err
	}

	return getLog, true, nil

}
