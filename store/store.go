package store

import (
	"database/sql"
	"log"
)

type Storer interface {
	Create()
	GetAll()
}

var db *sql.DB

func InitStore(sqlDb *sql.DB) {
	db = sqlDb
	log.Println("✅ Store")
}
