package siteTesting

import (
	"encoding/json"
	"net/http"
)

type PingResponse struct {
	Status string `json:"status"`
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PingResponse{Status: "ok"})
}

func StartServer(addr string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", PingHandler)
	return http.ListenAndServe(addr, mux)
}
