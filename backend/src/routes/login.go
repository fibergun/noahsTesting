package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type UserRequest struct {
	Username string `json:"user"`
}

func (s Server) login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering login")
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	group := r.PathValue("group")

	groupID, err := s.db.GetGroup(group)
	if err != nil {
		http.Error(w, "group not found", http.StatusNotFound)
		return
	}

	user := UserRequest{}
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userReturn, err := s.db.Login(strings.ToLower(user.Username), groupID)
	if err != nil {
		log.Println("getting user :", err)
		http.Error(w, fmt.Sprintf("getting user: %T", err), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(userReturn)
}
