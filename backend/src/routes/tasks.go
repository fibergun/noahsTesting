package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type TaskRequest struct {
	UserID int    `json:"userID"`
	Task   string `json:"task"`
}

type TasksRequest struct {
	UserID int `json:"userID"`
}

func (s Server) makeTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering makeTask")
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	task := TaskRequest{}
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userEntry, err := s.db.GetUserByID(task.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	taskEntry, err := s.db.InsertTask(task.Task, userEntry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("made task:", taskEntry)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(taskEntry)
}

func (s Server) getTasks(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		log.Println("method not allowed", r.Method)
		http.Error(w, "method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("userID"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tasks, err := s.db.GetAllTasksByUserID(userID)
	if err != nil {
		log.Println(err)

		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (s Server) getRandomTask(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		log.Println("method not allowed", r.Method)
		http.Error(w, "method not allowed.", http.StatusMethodNotAllowed)
		return
	}

	userID, err := strconv.Atoi(r.URL.Query().Get("userID"))
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := s.db.GetUserByID(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	randomTask, err := s.db.GetRandomTask(user.GroupID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	fmt.Println(randomTask)

	json.NewEncoder(w).Encode(randomTask)
}

//func (s Server)updateTask(w http.ResponseWriter, r *http.Request)  {
//
//}
