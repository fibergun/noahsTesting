package database

import (
	"database/sql"
	"fmt"
	"log"
)

func (db Database) InsertTask(task string, user UsersEntry) (TasksEntry, error) {

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

func (db Database) GetAllTasksByUserID(userID int) (AllTasks, error) {

	rows, err := db.Query("SELECT id, task, user_id, group_id, created_at FROM tasks WHERE user_id = ? ORDER BY id", userID)
	if err != nil {
		return AllTasks{}, err
	}
	defer rows.Close()

	getAllTasks := []TasksEntry{}
	for rows.Next() {
		var task TasksEntry
		err = rows.Scan(&task.ID, &task.Task, &task.UserID, &task.GroupID, &task.CreatedAt)
		if err != nil {
			return AllTasks{}, err
		}
		getAllTasks = append(getAllTasks, task)
	}

	points, err := db.GetPoints(userID)
	if err != nil {
		return AllTasks{}, err
	}

	return AllTasks{
		Tasks:  getAllTasks,
		Points: points,
	}, nil
}

func (db Database) GetRandomTask(groupID int64, userID int64) (TasksEntry, error) {
	getTask := &TasksEntry{}
	row := db.QueryRow(`SELECT id, task, group_id, user_id, created_at FROM tasks WHERE group_id = ? 
		AND id NOT IN (SELECT task_id FROM logs WHERE user_id = ?)
		ORDER BY RANDOM() LIMIT 1`, groupID, userID)

	err := row.Scan(&getTask.ID, &getTask.Task, &getTask.UserID, &getTask.GroupID, &getTask.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return TasksEntry{}, fmt.Errorf("no more available tasks for user")
		}
		log.Println("Error getting task: ", err)
		return TasksEntry{}, err
	}

	_, err = db.addTaskToLogs(getTask.ID, userID)
	if err != nil {
		return TasksEntry{}, fmt.Errorf("adding task to logs: %v", err)
	}

	return *getTask, nil

}

func (db Database) CompleteTask(taskID int) error {

	_, err := db.Exec("UPDATE logs SET completed = true WHERE task_id = ?", taskID)
	if err != nil {
		return err
	}
	return nil
}
