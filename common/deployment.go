package common

import (
	"github.com/google/go-github/github"
	"github.com/satori/go.uuid"
)

type DeploymentState int

const (
	DEPLOYMENT_PENDING DeploymentState = iota
	RUNNING_PRECONDITIONS
	RUNNING_PHASE
	DEPLOYMENT_DONE
	DEPLOYMENT_ERROR
)

func (state DeploymentState) String() string {
	switch state {
	case DEPLOYMENT_PENDING:
		return "DEPLOYMENT_PENDING"
	case RUNNING_PRECONDITIONS:
		return "RUNNING_PRECONDITIONS"
	case RUNNING_PHASE:
		return "RUNNING_PHASE"
	case DEPLOYMENT_DONE:
		return "DEPLOYMENT_DONE"
	case DEPLOYMENT_ERROR:
		return "DEPLOYMENT_ERROR"
	default:
		return "UNKNOWN"
	}
}

type DeploymentStatus struct {
	State DeploymentState
	Phase Phase
	Error error
}

type Deployment interface {
	UUID() uuid.UUID
	Application() Application
	Environment() string
	Ref() string
	GithubClient() *github.Client
	SetGithubClient(*github.Client)
	Flags() map[string]interface{}
	HasFlag(string) bool
	Flag(string) interface{}
	SetFlag(string, interface{})
	IsForce() bool

	Status() DeploymentStatus
}

type Application interface {
	Name() string
	Repository() Repository
	// Returns the strategy to use for deploying to a given environment or nil
	// if no strategy could be determined.
	StrategyForEnvironment(string) Strategy
}

type Repository interface {
	Owner() string
	Name() string
}

type Strategy interface {
	Preconditions() []Precondition
	Phases() []Phase
}

func CanonicalNameForRepository(repository Repository) string {
	owner := repository.Owner()
	name := repository.Name()
	return owner + "/" + name
}
