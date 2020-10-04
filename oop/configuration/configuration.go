package configuration

type configuration struct {
	githubAccessToken string
	workDir           string
}

func New(githubAccessToken string, workDir string) configuration {
	return configuration{githubAccessToken, workDir}
}

func (c configuration) GithubAccessToken() string {
	return c.githubAccessToken
}

func (c configuration) WorkDir() string {
	return c.workDir
}
