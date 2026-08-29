package routes

import (
	"log"
	"net/http"
	"noahsTesting/backend/src/database"
	"noahsTesting/backend/src/frontend"
)

type Server struct {
	server *http.ServeMux
	db     database.Database
}

func NewServer() Server {
	log.Println("Creating new server")

	return Server{
		server: http.NewServeMux(),
		db:     database.New(),
	}
}

func (s Server) StartServer(addr string) error {
	log.Println("Starting server...")
	s.loadRoutes()

	return http.ListenAndServe(addr, s.server)
}

func (s Server) loadRoutes() {
	log.Println("Loading routes...")

	s.server.HandleFunc("GET /api/ping", s.pingHandler)
	s.server.HandleFunc("POST /api/login/{group}", s.login)
	s.server.HandleFunc("POST /api/tasks/make", s.makeTask) //maybe also use to update?
	s.server.HandleFunc("GET /api/tasks/list", s.getTasks)
	s.server.HandleFunc("GET /api/tasks/random", s.getRandomTask)
	s.server.HandleFunc("POST /api/tasks/complete", s.completeTask)
	s.server.HandleFunc("GET /api/tasks/get", s.getTask)
	s.server.Handle("GET /", frontend.Handler())
}
