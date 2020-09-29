package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/apex/log"
	cliHandler "github.com/apex/log/handlers/cli"
	"github.com/dustin/go-humanize"
	"github.com/fatih/color"
	"github.com/google/go-github/github"
	"github.com/urfave/cli/v2"
	"github.com/wesleimp/github-mergetime/internal/merges"
	"golang.org/x/oauth2"
)

var (
	version = "v0.1.0"
)

func main() {
	log.SetHandler(cliHandler.Default)

	app := &cli.App{
		Name:      "github-mergetime",
		UsageText: "github-mergetime [options...] owner/repo",
		Usage:     "Lists the time it took a pull request to merge",
		Version:   version,
		Authors: []*cli.Author{{
			Name:  "Weslei Juan Novaes Pereira",
			Email: "wesleimsr@gmail.com",
		}},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "github-token",
				Aliases: []string{"t"},
				Usage:   "Github access token",
				Value:   "",
				EnvVars: []string{"GITHUB_TOKEN"},
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"V"},
				Usage:   "Enable verbose mode",
				Value:   false,
			},
			&cli.IntFlag{
				Name:  "page",
				Usage: "Page number",
				Value: 1,
			},
			&cli.IntFlag{
				Name:  "per-page",
				Usage: "Number of records per page",
				Value: 15,
			},
		},
		Action: run,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.WithError(err).Fatal("error executing cli")
	}
}

func run(c *cli.Context) error {
	if c.Bool("verbose") {
		log.SetLevel(log.DebugLevel)
	}

	if c.NArg() == 0 {
		return errors.New("github-mergetime requires exactly 1 argment")
	}
	arg := c.Args().Get(0)
	repo := strings.SplitN(arg, "/", 2)
	if len(repo) != 2 {
		return errors.New("Invalid argument. Must be like owner/repo")
	}

	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: c.String("github-token"),
	})
	tc := oauth2.NewClient(context.Background(), ts)

	options := []merges.Option{
		merges.WithClient(github.NewClient(tc)),
	}

	if c.IsSet("page") {
		options = append(options, merges.WithPage(c.Int("page")))
	}

	if c.IsSet("per-page") {
		options = append(options, merges.WithPerPage(c.Int("per-page")))
	}

	m := merges.New(options...)

	log.Debug("getting times")
	info, err := m.GetTimes(repo[0], repo[1])
	if err != nil {
		return err
	}

	log.Debug("listing times")
	for _, i := range info {
		title := fmt.Sprintf("%s - %s", color.New(color.Bold).Sprintf("#%d", i.Number), i.Title)
		log.Info(title)
		log.Info(humanize.RelTime(i.CreatedAt, i.MergedAt, "", ""))
		fmt.Println()
	}
	return nil
}
