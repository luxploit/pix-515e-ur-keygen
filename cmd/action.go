package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

func NewAction(actimpl Action) {
	flags := NewFlags(actimpl.Name())
	loadedActions[actimpl.Name()] = loadedAction{action: actimpl, flags: flags}
	actimpl.Init(flags)
}

func Usage() {
	exec, _ := os.Executable()
	exec = filepath.Base(exec)
	fmt.Printf("Usage: %s <action>\n", exec)

	for _, key := range loadedActions.SortedKeys() {
		loadedAction := loadedActions[key]
		fmt.Printf("  %s options:\n", loadedAction.action.Name())
		loadedAction.flags.Usage()
		fmt.Printf("  %s supported devices:\n", loadedAction.action.Name())
		loadedAction.action.Support()
		fmt.Println()
	}

	fmt.Printf("This help menu can be brought up again by running \"./%s help\"\n\n", exec)
	os.Exit(1)
}

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("No action specified!")
		Usage()
	}

	action := os.Args[1]
	if action == "help" {
		Usage()
	}

	act, ok := loadedActions[action]
	if !ok {
		fmt.Println("Invalid action specified!")
		Usage()
	}

	act.flags.Parse()
	res := act.action.Run()
	if res == Action_Failed {
		os.Exit(1)
	} else if res == Action_Invalid {
		Usage()
		os.Exit(1)
	}
}
