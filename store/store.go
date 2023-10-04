package store

import (
	"database/sql"
	"log"
)

var db *sql.DB

func InitStore(sqlDb *sql.DB) {
	db = sqlDb
	log.Println("✅ Store")
}
