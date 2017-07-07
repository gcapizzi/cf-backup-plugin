package main_test

import (
	. "code.cloudfoundry.org/backup-plugin"

	"io"

	"io/ioutil"

	"code.cloudfoundry.org/backup-plugin/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BackupPlugin", func() {
	var reader, writer = io.Pipe()
	var backupPlugin = BackupPlugin{Output: writer}

	Describe("GetMetadata", func() {
		It("Returns a PluginMetadata struct", func() {
			metadata := backupPlugin.GetMetadata()

			Expect(metadata.Name).To(Equal("BackupPlugin"))
		})
	})

	Describe("Run", func() {
		Describe("backup-service", func() {
			It("Sends a backup command to the service through the update-service CLI command", func() {
				cliConnection := new(fakes.FakeCliConnection)
				cliConnection.CliCommandReturnsOnCall(1, []string{"...", "Status: update in progress", "..."}, nil)
				cliConnection.CliCommandReturnsOnCall(2, []string{"...", "Status: update in progress", "..."}, nil)
				cliConnection.CliCommandReturnsOnCall(3, []string{"...", "Status: update succeeded", "Message: some message", "..."}, nil)

				go func() {
					backupPlugin.Run(cliConnection, []string{"backup-service", "service-name"})
					writer.Close()
				}()
				output, _ := ioutil.ReadAll(reader)

				Expect(cliConnection.CliCommandArgsForCall(0)).To(Equal([]string{"update-service", "service-name", "-c", `{"backup-action": "backup"}`}))

				Expect(string(output)).To(Equal("some message\n"))
			})
		})
	})
})
