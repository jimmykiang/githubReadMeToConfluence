package utils

import (
	"context"
	"fmt"
	"github.com/google/go-github/v49/github"
	"golang.org/x/oauth2"
	"log"
)

func FromGitHub(owner, repository, filePath, accessToken string) {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	readme, _, _, err := client.Repositories.GetContents(context.Background(), owner, repository, filePath, nil)
	if err != nil {
		log.Fatal(err)
	}
	readmeContent, _ := readme.GetContent()

	fmt.Println(readmeContent)
}
