package main

import (
	queue "ai-code-reviewer/queue"
	"fmt"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Println("Error loading .env")
	}

	router := gin.Default()

	router.POST("/webhook", handleWebhook)

	fmt.Println("Server running on port 8080")

	router.Run(":8080")
}

func handleWebhook(c *gin.Context) {

	var payload map[string]interface{}

	if err := c.BindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"error": "Invalid payload"})
		return
	}

	action := payload["action"]

	if action == "opened" || action == "synchronize" {

		repo := payload["repository"].(map[string]interface{})
		repoName := repo["name"].(string)

		ownerData := repo["owner"].(map[string]interface{})
		owner := ownerData["login"].(string)

		pr := payload["pull_request"].(map[string]interface{})
		prNumber := int(pr["number"].(float64))

		fmt.Println("Repo:", repoName)
		fmt.Println("Owner:", owner)
		fmt.Println("PR Number:", prNumber)

		job := queue.ReviewJob{
			Owner:    owner,
			Repo:     repoName,
			PRNumber: prNumber,
		}

		queue.PushReviewJob(job)

		fmt.Println("Review job added to queue")
	}

	c.JSON(200, gin.H{"status": "received"})
}
