package database

import (
	"database/sql"
	"log"
)

func SaveReview(
	db *sql.DB,
	prID int,
	reviewText string,
	score int,
) {

	query := `
	INSERT INTO reviews(pr_id, review_text, score)
	VALUES($1, $2, $3)
	`

	_, err := db.Exec(query, prID, reviewText, score)

	if err != nil {
		log.Println("Failed to save review:", err)
	}
}
