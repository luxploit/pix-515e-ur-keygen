package main

import (
	"luxploit.net/pixlic/agent"
	"luxploit.net/pixlic/cmd"
)

func main() {
	cmd.NewAction(&agent.PIX{})
	cmd.NewAction(&agent.ASA{})
	cmd.Execute()
}
