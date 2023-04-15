package create

import (
	"context"

	"github.com/google/go-github/github"
)

// set the website for the repo

func SetWebsite(client *github.Client, owner string, repo string, website string) (*github.Repository, error) {
	ctx := context.Background()

	Repository := &github.Repository{
		Homepage: &website,
	}

	_, _, err := client.Repositories.Edit(ctx, owner, repo, Repository)
	if err != nil {
		return nil, err
	}

	return Repository, nil

}

// set github checks for the repo

func SetChecks(client *github.Client, owner string, repo string, check bool) (*github.Repository, error) {
	ctx := context.Background()

	Repository := &github.Repository{
		HasIssues: &check,
	}

	_, _, err := client.Repositories.Edit(ctx, owner, repo, Repository)
	if err != nil {
		return nil, err
	}

	return Repository, nil

}

// update repo check status

func UpdateChecks(client *github.Client, owner string, repo string, check bool) (*github.Repository, error) {
	ctx := context.Background()

	Repository := &github.Repository{
		HasIssues: &check,
	}

	_, _, err := client.Repositories.Edit(ctx, owner, repo, Repository)
	if err != nil {
		return nil, err
	}

	return Repository, nil

}
