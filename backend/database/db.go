package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {

	connStr := "host=localhost user=khullar dbname=ai_code_reviewer sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Database not reachable")
	}

	return db
}
