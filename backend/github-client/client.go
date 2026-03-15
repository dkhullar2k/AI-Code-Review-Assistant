package githubclient

import (
	"context"
	"os"

	"github.com/google/go-github/v57/github"
	"golang.org/x/oauth2"
)

func NewGitHubClient() *github.Client {

	token := os.Getenv("GITHUB_TOKEN")

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(context.Background(), ts)

	client := github.NewClient(tc)

	return client
}
