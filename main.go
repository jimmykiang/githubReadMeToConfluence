package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/v49/github"
	"log"
)
import "golang.org/x/oauth2"

func main() {

	owner := "jimmykiang"
	repository := "testReadme"
	accessToken := "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"

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	readme, _, err := client.Repositories.GetReadme(context.Background(), owner, repository, nil)
	if err != nil {
		log.Fatal(err)
	}
	readmeContent, _ := readme.GetContent()
	fmt.Println(string(readmeContent))

	// ToDo below:

	// list all repositories for the authenticated user
	//repos, _, _ := client.Repositories.List(ctx, "", nil)
	//fmt.Println(repos)

	//// Load API keys
	//githubAPIKey := "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"
	////confluenceAPIKey := "YOUR_CONFLUENCE_API_KEY"
	//urlStr := "https://api.github.com/repos/jimmykiang/fluidengine/blob/main/README.md"
	//
	//// Initialize GitHub client
	//client := github.NewClient(nil)
	//req, err := client.NewRequest("GET", urlStr, nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//req.Header.Set("Authorization", "Token "+githubAPIKey)
	//readme, _, err := client.Repositories.GetReadme(context.Background(), "OWNER", "REPO", nil)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//readmeContent, _ := readme.GetContent()
	//fmt.Println(string(readmeContent))

	/*
		// Initialize Confluence client
		url := "https://YOUR_CONFLUENCE_URL/wiki/rest/api/content"
		req, err = http.NewRequest("POST", url, strings.NewReader(string(readmeContent)))
		if err != nil {
			log.Fatal(err)
		}
		req.SetBasicAuth("USERNAME", confluenceAPIKey)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))

	*/
}
