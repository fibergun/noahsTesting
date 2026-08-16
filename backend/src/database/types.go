package database

import (
	"database/sql"
	"time"
)

type Database struct {
	*sql.DB
}

type AllTasks struct {
	Tasks  []LogsEntry `json:"tasks"`
	Points int64       `json:"points"`
}

type UsersEntry struct {
	ID        int64     `json:"userID"`
	GroupID   int64     `json:"groupID"`
	Name      string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}

type TasksEntry struct {
	ID        int64     `json:"taskID"`
	Task      string    `json:"task"`
	GroupID   int64     `json:"groupID"`
	UserID    int64     `json:"userID"`
	CreatedAt time.Time `json:"createdAt"`
}

type LogsEntry struct {
	ID        int64      `json:"logID"`
	TaskID    int64      `json:"task"`
	UserID    int64      `json:"userID"`
	Completed bool       `json:"completed"`
	CreatedAt *time.Time `json:"createdAt"`
}
