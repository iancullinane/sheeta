package source

import (
	"context"

	"github.com/google/go-github/v44/github"
	"golang.org/x/oauth2"
)

func GithubClient(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return client
	// list all repositories for the authenticated user
	// repos, _, err := client.Repositories.List(ctx, "", nil)
}
