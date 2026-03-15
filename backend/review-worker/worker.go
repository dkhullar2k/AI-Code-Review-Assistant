package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/joho/godotenv"

	githubclient "ai-code-reviewer/github-client"
	queue "ai-code-reviewer/queue"
)

var ctx = context.Background()

func main() {

	err := godotenv.Load("../.env")

	if err != nil {
		fmt.Println("Error loading .env file")
	}

	fmt.Println("Review Worker Started")

	rdb := queue.NewRedisClient()

	for {

		result, err := rdb.BRPop(ctx, 0*time.Second, "review_jobs").Result()

		if err != nil {
			fmt.Println("Queue read error:", err)
			continue
		}

		jobData := result[1]

		var job queue.ReviewJob

		json.Unmarshal([]byte(jobData), &job)

		fmt.Println("Processing job:", job)

		githubclient.GetPullRequestFiles(
			job.Owner,
			job.Repo,
			job.PRNumber,
		)
	}
}
