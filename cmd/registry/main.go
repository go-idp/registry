package main

import (
	"github.com/go-idp/registry"
	"github.com/go-idp/registry/cmd/registry/commands"
	"github.com/go-zoox/cli"
)

func main() {
	app := cli.NewMultipleProgram(&cli.MultipleProgramConfig{
		Name:    "registry",
		Usage:   "Registry Service for IDP",
		Version: registry.Version,
	})

	// server
	commands.RegisterServer(app)
	// client
	// commands.RegisterClient(app)

	app.Run()
}
