package repo

import (
	"context"
	"fmt"
	"github.com/google/go-github/v52/github"
	"golang.org/x/oauth2"
	"log"
	"os"
)

type Repo interface {
	Create(name string)
}

type repo struct {
	Client *github.Client
}

func (r *repo) Create(name string) {
	ctx := context.Background()
	org := os.Getenv("MINI_PLATFORM_ORG")
	repository, response, err := r.Client.Repositories.Create(ctx, org, &github.Repository{
		Name: &name,
	})

	if err != nil {
		log.Fatalln("error while creating repository", err.Error())
	}

	if response.StatusCode == 201 {
		log.Println(fmt.Sprintf(`
Hey! Your repository is ready
Get the repository: git clone %s`, *repository.CloneURL))
	} else {
		log.Fatalln("Sorry! was not to possible creating your repository")
	}
}

func NewRepo() Repo {
	personalAccessToken := os.Getenv("MINI_PLATFORM_GH_TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: personalAccessToken})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return &repo{
		Client: client,
	}
}
