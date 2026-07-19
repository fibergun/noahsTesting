package routes

import (
	"encoding/json"
	"net/http"
)

type PingResponse struct {
	Status string `json:"status"`
}

func (s Server) pingHandler(w http.ResponseWriter, r *http.Request) {
	//log.Println("PingHandler")
	//cookie, err := r.Cookie("session")
	//if err != nil {
	//	log.Println("getting cookie: ", err)
	//}
	//fmt.Println(cookie.Value)

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(PingResponse{Status: "ok"})
}
