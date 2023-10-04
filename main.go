package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevinferri/tasky-go/config"
	"github.com/kevinferri/tasky-go/db"
	"github.com/kevinferri/tasky-go/handlers"
	"github.com/kevinferri/tasky-go/middleware"
	"github.com/kevinferri/tasky-go/store"
)

type Server struct {
	Port   string
	Router *mux.Router
	DB     *sql.DB
}

func (s *Server) Start() {
	handlers.InitHandlers(s.Router)
	middleware.InitMiddleware(s.Router)
	store.InitStore(s.DB)

	log.Println("🚀", "Server running on port", s.Port)
	log.Fatal(http.ListenAndServe("localhost:"+s.Port, s.Router))
}

func main() {
	config.Init()

	server := &Server{
		Port:   config.GetEnv("PORT"),
		Router: mux.NewRouter().StrictSlash(true),
		DB:     db.InitDB(),
	}

	server.Start()
}
