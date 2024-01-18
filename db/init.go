package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "ranvijay"
	password = ""
	dbname   = "tasks"
)

func Init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the database")

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS tasks1 (
			id VARCHAR(36) PRIMARY KEY,
			title VARCHAR(255),
			description TEXT,
			priority INT,
			time BIGINT
		)
	`)
	if err != nil {
		fmt.Println("Error creating tasks table:", err)
	}

	fmt.Println("Table checked.")
}
