package main

import (
	"io"

	"os"

	"fmt"

	"strings"

	"code.cloudfoundry.org/cli/plugin"
)

type BackupPlugin struct {
	Output io.Writer
}

var actions = map[string]string{
	"backup-service":       "backup",
	"restore-service":      "restore",
	"list-service-backups": "list",
}

func (c *BackupPlugin) Run(cliConnection plugin.CliConnection, args []string) {
	serviceName := args[1]
	action := actions[args[0]]

	cliConnection.CliCommand("update-service", serviceName, "-c", `{"backup-action": "`+action+`"}`)

	for {
		serviceInfoOutput, _ := cliConnection.CliCommand("service", serviceName)
		if anyLineContains(serviceInfoOutput, "Status: update succeeded") {
			message := extractMessageFrom(serviceInfoOutput)
			fmt.Fprintln(c.Output, message)
			return
		}
	}
}

func anyLineContains(lines []string, substring string) bool {
	for _, line := range lines {
		if strings.Contains(line, substring) {
			return true
		}
	}
	return false
}

func extractMessageFrom(lines []string) string {
	for _, line := range lines {
		if strings.Contains(line, "Message:") {
			return strings.Split(line, ": ")[1]
		}
	}
	return ""
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
				UsageDetails: plugin.Usage{
					Usage: "backup-service\n   cf backup-service [service-name]",
				},
			},
			{
				Name:     "restore-service",
				HelpText: "Restore a service",
				UsageDetails: plugin.Usage{
					Usage: "restore-service\n   cf restore-service [service-name]",
				},
			},
			{
				Name:     "list-service-backups",
				HelpText: "List all backups for a service",
				UsageDetails: plugin.Usage{
					Usage: "list-service-backups\n   cf list-service-backups [service-name]",
				},
			},
		},
	}
}

func main() {
	plugin.Start(&BackupPlugin{Output: os.Stdout})
}
