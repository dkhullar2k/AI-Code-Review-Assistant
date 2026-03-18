package main

import (
	database "ai-code-reviewer/database"

	"github.com/gin-gonic/gin"

	"time"

	"github.com/gin-contrib/cors"
)

func main() {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/reviews", getReviews)

	router.Run(":8090")
}

func getReviews(c *gin.Context) {

	db := database.ConnectDB()

	rows, err := db.Query(`
		SELECT pr_id, review_text, score, created_at
		FROM reviews
		ORDER BY created_at DESC
	`)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	type Review struct {
		PRID      int    `json:"pr_id"`
		Text      string `json:"review_text"`
		Score     int    `json:"score"`
		CreatedAt string `json:"created_at"`
	}

	reviews := []Review{}

	for rows.Next() {

		var r Review

		rows.Scan(&r.PRID, &r.Text, &r.Score, &r.CreatedAt)

		reviews = append(reviews, r)
	}

	c.JSON(200, reviews)
}
