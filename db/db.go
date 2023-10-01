package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kevinferri/tasky-go/config"
	_ "github.com/lib/pq"
)

type DBConnection struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func (dbc *DBConnection) getConnectionString() string {

	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s",
		dbc.User, dbc.Password, dbc.Host, dbc.Port, dbc.DBName)

}

func InitDB() *sql.DB {
	dbConn := DBConnection{
		Host:     config.GetEnv("POSTGRES_URL"),
		Port:     config.GetEnv("POSTGRES_PORT"),
		User:     config.GetEnv("POSTGRES_USER"),
		Password: config.GetEnv("POSTGRES_PASSWORD"),
		DBName:   config.GetEnv("POSTGRES_DB"),
	}

	db, err := sql.Open("postgres", dbConn.getConnectionString())

	if err != nil {
		log.Fatal("Could not connect to the DB:", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Could not ping database:", err.Error())
	}

	log.Println("Connected to postgres database successfully")

	return db
}
