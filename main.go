package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

import utils "jimmykiang/githubReadMeToConfluence/utils"

func main() {

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "owner",
				Aliases:  []string{"o"},
				Usage:    "Owner of the GitHub Repository.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "repository",
				Aliases:  []string{"r"},
				Usage:    "Repository name.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "filepath",
				Aliases:  []string{"f"},
				Usage:    "File including its path from the repository root.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "githubtoken",
				Aliases:  []string{"gt"},
				Usage:    "Valid Access Token for GitHub.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "confluencelocation",
				Aliases:  []string{"cl"},
				Usage:    "Domain of the Confluence owner. ",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "userName",
				Aliases:  []string{"u"},
				Usage:    "Valid username for Confluence.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "confluencetoken",
				Aliases:  []string{"ct"},
				Usage:    "Valid Access Token for Confluence.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "title",
				Aliases:  []string{"t"},
				Usage:    "An initial one-word title for the Confluence page.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "key",
				Aliases:  []string{"k"},
				Usage:    "Key for the Confluence Space.",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "id",
				Aliases: []string{"i"},
				Usage:   "Id for creating a subPage (Leaving it blank will create the page at the root of the Confluence space tree.)",
			},
		},
		Action: func(cCtx *cli.Context) error {
			owner := cCtx.String("owner")
			repository := cCtx.String("repository")
			filepath := cCtx.String("filepath")
			githubtoken := cCtx.String("githubtoken")
			userName := cCtx.String("userName")
			confluenceToken := cCtx.String("confluencetoken")
			title := cCtx.String("title")
			key := cCtx.String("key")
			id := cCtx.String("id")

			confluencelocation := cCtx.String("confluencelocation")
			confluencelocation = strings.TrimSpace(confluencelocation)

			payLoad := utils.FromGitHub(owner, repository, filepath, githubtoken)

			//re := regexp.MustCompile(`&`)
			//payLoad = re.ReplaceAllString(payLoad, "&amp;")

			// remove i.e.: [Add pipeline permissions to a repository resource]
			//re = regexp.MustCompile(`\[(.*?)\]`)
			//payLoad = re.ReplaceAllString(payLoad, "")

			html := utils.MarkDownToHtml(payLoad)
			utils.ToConfluence(string(html), confluencelocation, userName, confluenceToken, title, key, id)
			return nil
		},

		Version: "v1.0 2023",
	}

	app.CustomAppHelpTemplate = `
			***	GitHub To Confluence tool {{.Version}} ***

						A product from: 
			*** Way of Working and Practices ***

A CLI convenience tool meant to help you replicate the content from your GitHub file
into your favourite Confluence page. One at a time!

Disclaimer:
This software is provided "as is," without warranty of any kind.
The user assumes full responsibility arising from its use.


USAGE:
   {{.HelpName}} {{if .VisibleFlags}}[global options] [values]{{end}}{{if .Commands}}{{end}}
   {{if len .Authors}}
AUTHOR:
   {{range .Authors}}{{ . }}{{end}}
   {{end}}{{if .Commands}}
GLOBAL OPTIONS:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{if .Copyright }}
COPYRIGHT:
   {{.Copyright}}
   {{end}}{{if .Version}}
   {{end}}
`

	if err := app.Run(os.Args); err != nil {
		fmt.Println("Something is missing:", err)
	}

}
