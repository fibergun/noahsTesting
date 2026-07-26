package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
}

type User struct {
	Username string `json:"user"`
	userID   int64  `json:"userid"`
}

func (s Server) login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering login")
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	user.userID = s.db.Login(user.Username)

	fmt.Printf("user: %v\n", user)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(user)
}
