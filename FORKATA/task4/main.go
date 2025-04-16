package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v53/github"
	"golang.org/x/oauth2"
	"log"
)

type Item struct {
	Title    string
	Описание string
	Link     string
}

type GithubLister interface {
	GetItems(ctx context.Context, username string) ([]Item, error)
}

type GeneralGithubLister interface {
	GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error)
}

type GithubGist struct{
	client *github.Client
}

func NewGithubGist(c *github.Client) *GithubGist{
	return &GithubGist{
		client: c,
	}
}

func (gg *GithubGist) GetItems(ctx context.Context, username string) ([]Item, error){
	opt := &github.GistListOptions{ListOptions: github.ListOptions{PerPage: 1000}}

	body, _, err := gg.client.Gists.List(ctx, username, opt)
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
		items = append(items, Item{Title: title, Описание: desc, Link: url})
	}
	return items, err
}

type GithubRepo struct{
	client *github.Client
}


func NewGithubRepo(c *github.Client) *GithubRepo{
	return &GithubRepo{
		client: c,
	}
}

func (gg *GithubRepo) GetItems(ctx context.Context, username string) ([]Item, error){
	opt := &github.RepositoryListOptions{ListOptions: github.ListOptions{PerPage: 1000}}

	body, _, err := gg.client.Repositories.List(ctx, username, opt)
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
		items = append(items, Item{Title: title, Описание: desc, Link: url})
	}
	return items, err
}

type GeneralGithub struct{
	client *github.Client
}

func NewGeneralGithub(c *github.Client) *GeneralGithub{
	return &GeneralGithub{
		client: c,
	}
}

func (gg *GeneralGithub) GetItems(ctx context.Context, username string, strategy GithubLister) ([]Item, error){
	return strategy.GetItems(ctx, username)
}

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "your-access-token"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	gist := NewGithubGist(client)
	repo := NewGithubRepo(client)

	gg := NewGeneralGithub(client)

	data, err := gg.GetItems(context.Background(), "ptflp", gist)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)

	data, err = gg.GetItems(context.Background(), "ptflp", repo)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
}
