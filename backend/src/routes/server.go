package routes

import (
	"net/http"
)

type Server struct {
	*http.ServeMux
}

func NewServer() Server {
	server := http.NewServeMux()

	return Server{server}
}

func (s Server) StartServer(addr string) error {
	s.loadRoutes()

	return http.ListenAndServe(addr, s)
}

func (s Server) loadRoutes() {
	s.HandleFunc("/ping", s.pingHandler)
	s.HandleFunc("/user/{userId}", s.login)
}
