package utils

import (
	"context"
	"fmt"
	"github.com/google/go-github/v49/github"
	"github.com/russross/blackfriday/v2"
	confluence "github.com/virtomize/confluence-go-api"
	"golang.org/x/oauth2"
	"log"
)

func FromGitHub(owner, repository, filePath, accessToken string) string {

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

	return readmeContent
}

func ToConfluence(payLoad, confluencelocation, userName, token, title, key, id string) {

	confluencelocation = "https://" + confluencelocation + "/wiki/rest/api"

	client, err := confluence.NewAPI(confluencelocation, userName, token)
	if err != nil {
		log.Fatal("Something went wrong: ", err)
	}

	var content *confluence.Content

	if id == "" {

		content = &confluence.Content{
			Type:  "page", // can also be blogpost
			Title: title,  // page title
			Body: confluence.Body{
				Storage: confluence.Storage{
					Value:          payLoad, // your page content here.
					Representation: "storage",
				},
			},
			Version: &confluence.Version{
				Number: 1,
			},
			Space: &confluence.Space{
				Key: key, // Space
			},
		}

	} else {

		// Define the content of the new page
		content = &confluence.Content{
			Type:  "page", // can also be blogpost
			Title: title,  // page title
			Ancestors: []confluence.Ancestor{
				{
					ID: id, // ancestor-id optional if you want to create sub-pages
				},
			},
			Body: confluence.Body{
				Storage: confluence.Storage{
					Value:          payLoad, // your page content here.
					Representation: "storage",
				},
			},
			Version: &confluence.Version{
				Number: 1,
			},
			Space: &confluence.Space{
				Key: key, // Space
			},
		}
	}

	c, err := client.CreateContent(content)
	if err != nil {
		log.Fatal("Something went wrong: ", err)
	}

	fmt.Printf("%+v\n", c)
}

func MarkDownToHtml(payLoad string) []byte {

	markdown := []byte(payLoad)

	html := blackfriday.Run(markdown)

	fmt.Println(string(html))
	return html
}
