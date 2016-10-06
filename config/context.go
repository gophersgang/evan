package config

import (
	"github.com/google/go-github/github"
)

type Deployment interface {
	Application() *Application
	Ref() string
	GithubClient() *github.Client
	Flags() map[string]interface{}

	// Some object representing the request that initiated this deployment,
	// eg. `*github.DeploymentEvent`.
	Initiator() interface{}
}
