package common

import "github.com/InkShaStudio/go-command"

var manifest *PluginManifest

func LoadManifest(data *PluginManifest) {
	manifest = data
}

func WithCommand(cmd *command.SCommand) *command.SCommand {
	cmd.AddSubCommand(version())

	return cmd
}
