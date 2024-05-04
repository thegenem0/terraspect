package config

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

func MasterUser(scope constructs.Construct) string {
	masterUser := "terraspect_admin"

	ctxValue := scope.Node().TryGetContext(jsii.String("masterUser"))
	if v, ok := ctxValue.(string); ok {
		masterUser = v
	}

	return masterUser
}
