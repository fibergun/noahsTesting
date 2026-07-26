package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Status string `json:"status"`
}

type User struct {
	Username string `json:"user"`
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

	strings.ToLower(user.Username)

	userReturn, err := s.db.Login(strings.ToLower(user.Username))
	if err != nil {
		log.Print("User login error", err)
	}

	fmt.Printf("user: %v\n", user)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(userReturn)
}
