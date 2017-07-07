package main

import (
	"fmt"

	"code.cloudfoundry.org/cli/plugin"
)

type BackupPlugin struct{}

func (c *BackupPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] == "backup-service" {
		fmt.Println("Running the backup-service command")
	}
}

func (c *BackupPlugin) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "BackupPlugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 0,
			Build: 1,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 7,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "backup-service",
				HelpText: "Backup a service",

				// UsageDetails is optional
				// It is used to show help of usage of each command
				UsageDetails: plugin.Usage{
					Usage: "backup-service\n   cf backup-service [service-name]",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(BackupPlugin))
}
