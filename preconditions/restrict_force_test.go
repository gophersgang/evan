package preconditions

import (
	"testing"

	"github.com/Everlane/evan/common"
	"github.com/Everlane/evan/context"

	"github.com/stretchr/testify/require"
)

func newDeployment(environment string, flags map[string]interface{}) *context.Deployment {
	var app common.Application = nil
	var strategy common.Strategy = nil
	ref := ""
	return context.NewDeployment(app, environment, strategy, ref, flags)
}

func TestIgnoresIfNotForce(t *testing.T) {
	deployment := newDeployment("development", map[string]interface{}{
		"force": false,
	})

	precondition := &RestrictForcePrecondition{Safelist: []string{}}
	require.Nil(t, precondition.Status(deployment))
}

func TestPassesIfOnSafelist(t *testing.T) {
	deployment := newDeployment("development", map[string]interface{}{
		"force": true,
	})

	precondition := &RestrictForcePrecondition{Safelist: []string{"development"}}
	require.Nil(t, precondition.Status(deployment))
}

func TestErrorsIfNotOnSafelist(t *testing.T) {
	deployment := newDeployment("production", map[string]interface{}{
		"force": true,
	})

	precondition := &RestrictForcePrecondition{Safelist: []string{"development"}}
	require.NotNil(t, precondition.Status(deployment))
}
