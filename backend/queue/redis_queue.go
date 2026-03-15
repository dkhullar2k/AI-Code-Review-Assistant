package queue

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func NewRedisClient() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return rdb
}

type ReviewJob struct {
	Owner    string
	Repo     string
	PRNumber int
}

func PushReviewJob(job ReviewJob) error {

	rdb := NewRedisClient()

	data, _ := json.Marshal(job)

	return rdb.LPush(ctx, "review_jobs", data).Err()
}
