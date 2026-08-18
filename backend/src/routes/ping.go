package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

type PingResponse struct {
	Status string `json:"status"`
}

func (s Server) pingHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("pingHandler")
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PingResponse{Status: "ok"})
}
