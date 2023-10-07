package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kevinferri/tasky-go/config"
	"github.com/kevinferri/tasky-go/db"
	"github.com/kevinferri/tasky-go/handlers"
	"github.com/kevinferri/tasky-go/middleware"
)

func main() {
	config.Init()

	db := db.InitDB()
	port := config.GetEnv("PORT")
	router := mux.NewRouter().StrictSlash(true)

	middleware.InitMiddleware(router)
	handlers.InitHandlers(handlers.NewHandler(router, db))

	log.Println("🚀", "Server running on port", port)
	log.Fatal(http.ListenAndServe("localhost:"+port, router))
}
