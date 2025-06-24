package cmd

import (
	"sort"
)

type ActionStatus int

const (
	Action_Success = ActionStatus(0)
	Action_Failed  = ActionStatus(1)
	Action_Invalid = ActionStatus(2)
)

type loadedAction struct {
	action Action
	flags  *Flags
}

type Action interface {
	Name() string
	Init(*Flags)
	Support()
	Run() ActionStatus
}

type actionMap map[string]loadedAction

var loadedActions = actionMap{}

func (a actionMap) SortedKeys() []string {
	keys := make([]string, 0, len(a))
	for k := range a {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
