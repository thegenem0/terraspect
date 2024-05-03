package change

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/thegenem0/terraspect_server/pkg/reflector"
)

func BuildChanges(changeData []*tfjson.ResourceChange) []Change {
	ref := reflector.NewReflectorModule()
	resources := make([]Change, 0)

	for _, change := range changeData {
		if change.Change.Actions.NoOp() {
			continue
		} else {
			resources = append(resources, addChangeResource(
				change.Change.Actions,
				change.Address,
				change.PreviousAddress,
				ref.HandleChanges(change.Change.Before, change.Change.After),
			))
		}
	}

	return resources
}

func addChangeResource(
	actions tfjson.Actions,
	address string,
	previousAddress string,
	changes reflector.ChangeData,
) Change {
	change := ChangeItem{
		Actions:         actions,
		Address:         address,
		PreviousAddress: previousAddress,
		Changes:         changes,
	}

	return Change{
		ModKey:  address,
		Changes: []ChangeItem{change},
	}
}
