package main

import (
	"context"
	"testing"

	"github.com/google/go-github/v53/github"
	"github.com/stretchr/testify/assert"
)

type MockRepoLister struct{}
type MockGistLister struct{}


func (mock *MockRepoLister) List(ctx context.Context, username string, opt *github.RepositoryListOptions) ([]*github.Repository, *github.Response, error){
	repo := []*github.Repository{
		{Name: github.String("Repository 1"), Description: github.String("Desc 1"), HTMLURL: github.String("example1@mail.com")},
		{Name: github.String("Repository 2"), Description: github.String("Desc 2"), HTMLURL: github.String("example2@mail.com")},
		{Name: github.String("Repository 3"), Description: github.String("Desc 3"), HTMLURL: github.String("example3@mail.com")},
	}
	return repo, nil, nil
}

func (mock *MockGistLister) List(ctx context.Context, username string, opt *github.GistListOptions) ([]*github.Gist, *github.Response, error){
	gitHubRep := []*github.Gist{
		{Description: github.String("Desc 1"), HTMLURL: github.String("example1@mail.com")},
		{Description: github.String("Desc 2"), HTMLURL: github.String("example2@mail.com")},
		{Description: github.String("Desc 3"), HTMLURL: github.String("example3@mail.com")},
	}
	return gitHubRep, nil, nil
}

func TestMock(t *testing.T){
	a := make(map[string][]Item)
	g := &GithubAdapter{
		RepoList: &MockRepoLister{},
		GistList: &MockGistLister{},
	}

	gP := &GithubProxy{
		github: g,
		cache: a,
	}


	TruebodyGist := []Item{{"Desc 1", "TASK: Desc 1", "example1@mail.com"}, {"Desc 2", "TASK: Desc 2", "example2@mail.com"}, {"Desc 3", "TASK: Desc 3", "example3@mail.com"}}
	TruebodyRep := []Item{{"Repository 1", "Desc 1", "example1@mail.com"}, {"Repository 2", "Desc 2", "example2@mail.com"}, {"Repository 3", "Desc 3", "example3@mail.com"}}
	bodyGist, err1 := gP.GetGists(context.Background(), "ptflp")
	bodyRepo, err2 := gP.GetRepos(context.Background(), "ptflp")

	assert.Equal(t, bodyGist, TruebodyGist)
	assert.Equal(t, bodyRepo, TruebodyRep)
	assert.NoError(t, err1)
	assert.NoError(t, err2)
}