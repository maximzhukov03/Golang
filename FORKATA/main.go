package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		// &oauth2.Token{AccessToken: ""},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	// g := &GithubAdapter{
	// 	RepoList: &MockRepoLister{},
	// 	GistList: &MockGistLister{},
	// }
	g := NewGithubAdapter(client)

	fmt.Println(g.GetGists(context.Background(), "ptflp"))
	fmt.Println(g.GetRepos(context.Background(), "ptflp"))
}

type RepoLister interface {
	List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error)
}

type GistLister interface {
	List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error)
}

type Githuber interface {
	GetGists(ctx context.Context, username string) ([]Item, error)
	GetRepos(ctx context.Context, username string) ([]Item, error) // opt := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 1000}}
}

type GithubAdapter struct {
	RepoList RepoLister
	GistList GistLister
}

func NewGithubAdapter(githubClient *github.Client) *GithubAdapter {
	g := &GithubAdapter{
		RepoList: githubClient.Repositories,
		GistList: githubClient.Gists,
	}

	return g
}

func (gitAd *GithubAdapter) GetGists(ctx context.Context, username string) ([]Item, error) {
	opt := &github.GistListOptions{}

	body, _, err := gitAd.GistList.List(ctx, username, opt)
	if err != nil {
		return nil, err
	}
	var items []Item
	for _, elem := range body {
		title := ""
		if elem.Description != nil {
			title = *elem.Description
		}
		url := ""
		if elem.HTMLURL != nil {
			url = *elem.HTMLURL
		}
		desc := "TASK: " + title
		items = append(items, Item{Title: title, Description: desc, Link: url})
	}
	return items, err
}

func (gitAd *GithubAdapter) GetRepos(ctx context.Context, username string) ([]Item, error) {
	opt := &github.RepositoryListOptions{}

	body, _, err := gitAd.RepoList.List(ctx, username, opt)
	if err != nil {
		return nil, err
	}
	var items []Item
	for _, elem := range body {
		title := ""
		if elem.Name != nil {
			title = *elem.Name
		}
		desc := ""
		if elem.Description != nil {
			desc = *elem.Description
		}
		url := ""
		if elem.HTMLURL != nil {
			url = *elem.HTMLURL
		}
		items = append(items, Item{Title: title, Description: desc, Link: url})
	}
	return items, err
}

type Item struct {
	Title       string
	Description string
	Link        string
}
