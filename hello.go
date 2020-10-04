package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	githubAccessToken := viper.GetString("GITHUB_ACCESS_TOKEN")
	workDir := viper.GetString("WORKDIR")

	workDir = workDir

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: githubAccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	//orgs, _, err := client.Organizations.List(ctx, "medined", nil)
	client.Organizations.List(ctx, "medined", nil)

	//Info("git clone https://github.com/go-git/go-git")
	//_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
	//	URL:      "https://github.com/go-git/go-git",
	//	Progress: os.Stdout,
	//})
	//CheckIfError(err)
}
