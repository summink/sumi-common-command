package common

import (
	"fmt"

	"github.com/InkShaStudio/go-command"
)

func version() *command.SCommand {
	cmd := command.
		NewCommand("version").
		ChangeDescription("see version").
		RegisterHandler(func(cmd *command.SCommand) {
			println(fmt.Sprintf("%s %s", manifest.Name, manifest.Version))
		})

	return cmd
}
