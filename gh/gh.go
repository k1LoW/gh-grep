package gh

import (
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"
	"github.com/k1LoW/go-github-client/v53/factory"
)

type Gh struct {
	client *github.Client
}

func New() (*Gh, error) {
	client, err := factory.NewGithubClient()
	if err != nil {
		return nil, err
	}
	return &Gh{
		client: client,
	}, nil
}

func (g *Gh) Client() *github.Client {
	return g.client
}

func (g *Gh) Repositories(ctx context.Context, owner string) ([]string, error) {
	repos := []string{}

	u, _, err := g.client.Users.Get(ctx, owner)
	if err != nil {
		return nil, err
	}
	if u.GetType() == "User" {
		// User
		page := 1
		for {
			rs, res, err := g.client.Repositories.List(ctx, owner, &github.RepositoryListOptions{
				ListOptions: github.ListOptions{
					Page:    page,
					PerPage: 100,
				},
			})
			if err != nil {
				return nil, err
			}
			for _, r := range rs {
				repos = append(repos, *r.Name)
			}
			if res.NextPage == 0 {
				break
			}
			page = res.NextPage
		}
	} else {
		// Organization
		page := 1
		for {
			rs, res, err := g.client.Repositories.ListByOrg(ctx, owner, &github.RepositoryListByOrgOptions{
				ListOptions: github.ListOptions{
					Page:    page,
					PerPage: 100,
				},
			})
			if err != nil {
				return nil, err
			}
			for _, r := range rs {
				repos = append(repos, *r.Name)
			}
			if res.NextPage == 0 {
				break
			}
			page = res.NextPage
		}
	}

	return repos, nil
}

func (g *Gh) ContentURL(ctx context.Context, owner, repo, path string) (string, error) {
	fc, _, _, err := g.client.Repositories.GetContents(ctx, owner, repo, path, &github.RepositoryContentGetOptions{})
	if err != nil {
		return "", err
	}
	if fc == nil {
		return "", fmt.Errorf("%s is not file", path)
	}
	return fc.GetHTMLURL(), nil
}
