package routes

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
}

type User struct {
	Username string `json:"username"`
}

func (s Server) login(w http.ResponseWriter, r *http.Request) {

	userId := r.PathValue("userId")

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PingResponse{Status: userId})
}
