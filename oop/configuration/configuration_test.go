package configuration

import (
	"fmt"
	"testing"
)

func init() {
	fmt.Println("aaa")
}

func TestGithubAccessToken(t *testing.T) {
	githubAccessToken := "AAA"
	workDir := "BBB"
	c := New(githubAccessToken, workDir)

	want := githubAccessToken
	if got := c.GithubAccessToken(); got != want {
		t.Errorf("Got: %q, Wanted: %q", got, want)
	}
}

func TestWorkDir(t *testing.T) {
	githubAccessToken := "AAA"
	workDir := "BBB"
	c := New(githubAccessToken, workDir)

	want := workDir
	if got := c.WorkDir(); got != want {
		t.Errorf("Got: %q, Wanted: %q", got, want)
	}
}
