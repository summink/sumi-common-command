package common

import "github.com/InkShaStudio/go-command"

var manifest *PluginManifest

func init() {
	var manifestFile = "manifest.json"

	data, err := LoadManifestByFile(manifestFile)

	if err != nil {
		panic(err)
	}

	manifest = data
}

func WithCommand(cmd *command.SCommand) *command.SCommand {
	cmd.AddSubCommand(version())

	return cmd
}
