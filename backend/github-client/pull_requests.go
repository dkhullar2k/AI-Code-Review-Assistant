package githubclient

import (
	aireview "ai-code-reviewer/ai-review"
	"ai-code-reviewer/database"
	"context"
	"fmt"

	"github.com/google/go-github/v57/github"
)

func GetPullRequestFiles(owner string, repo string, prNumber int) {

	client := NewGitHubClient()

	ctx := context.Background()

	files, _, err := client.PullRequests.ListFiles(
		ctx,
		owner,
		repo,
		prNumber,
		nil,
	)

	if err != nil {
		fmt.Println("Error fetching PR files:", err)
		return
	}

	for _, file := range files {

		fmt.Println("File:", file.GetFilename())

		diff := file.GetPatch()

		if diff != "" {

			fmt.Println("Sending diff to AI...")

			review := aireview.ReviewCode(diff)

			db := database.ConnectDB()

			database.SaveReview(
				db,
				prNumber,
				review.ReviewMarkdown,
				review.Score,
			)

			comment := "🤖 **AI Code Review**\n\n" + review.ReviewMarkdown

			CreatePRComment(owner, repo, prNumber, comment)
		}
	}
}

func CreatePRComment(owner string, repo string, prNumber int, comment string) {

	client := NewGitHubClient()

	ctx := context.Background()

	prComment := &github.IssueComment{
		Body: github.String(comment),
	}

	_, _, err := client.Issues.CreateComment(
		ctx,
		owner,
		repo,
		prNumber,
		prComment,
	)

	if err != nil {
		fmt.Println("Failed to create comment:", err)
		return
	}

	fmt.Println("PR comment posted successfully")
}
