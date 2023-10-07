package store

import (
	"log"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitStore(sqlDb *sqlx.DB) {
	db = sqlDb
	log.Println("✅ Store")
}
