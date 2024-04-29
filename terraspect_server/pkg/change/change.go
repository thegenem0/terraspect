package change

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/thegenem0/terraspect_server/pkg/reflector"
)

type IChangeModule interface {
	GetResourceKeys() []ResourceKey
	GetChanges() []Change
	BuildChanges(changeData []*tfjson.ResourceChange)
	IsValidKey(modKey string, key string) bool
}

type ChangeModule struct {
	reflectorService  reflector.IReflectorModule
	validResourceKeys []ResourceKey
	changes           []Change
}

func NewChangeModule(reflectorService reflector.IReflectorModule) *ChangeModule {
	return &ChangeModule{
		reflectorService:  reflectorService,
		validResourceKeys: make([]ResourceKey, 0),
		changes:           make([]Change, 0),
	}
}

func (cs *ChangeModule) GetResourceKeys() []ResourceKey {
	return cs.validResourceKeys
}

func (cs *ChangeModule) IsValidKey(modKey string, key string) bool {
	for _, k := range cs.validResourceKeys {
		if k.ModKey == modKey {
			for _, v := range k.Keys {
				if v == key {
					return true
				}
			}
		}
	}
	return false
}

func (cs *ChangeModule) GetChanges() []Change {
	return cs.changes
}

func (cs *ChangeModule) BuildChanges(changeData []*tfjson.ResourceChange) {
	for _, change := range changeData {
		if change.Change.Actions.NoOp() {
			continue
		} else {
			cs.addChangeResource(
				change.Change.Actions,
				change.Address,
				change.PreviousAddress,
				cs.reflectorService.HandleChanges(change.Change.Before, change.Change.After),
			)
		}
	}
}

func (cs *ChangeModule) addChangeResource(
	actions tfjson.Actions,
	address string,
	previousAddress string,
	changes reflector.ChangeData,
) {
	change := ChangeItem{
		Actions:         actions,
		Address:         address,
		PreviousAddress: previousAddress,
		Changes:         changes,
	}

	cs.changes = append(cs.changes, Change{
		ModKey:  address,
		Changes: []ChangeItem{change},
	})
}
