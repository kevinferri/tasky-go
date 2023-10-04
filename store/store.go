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

func idk() {
	var snippet Snippet

	query := `select * from snippets WHERE id=$1;`
	row := db.QueryRow(query, id)
	err := row.Scan(&snippet.ID, &snippet.Title)

	return snippet, err
}
