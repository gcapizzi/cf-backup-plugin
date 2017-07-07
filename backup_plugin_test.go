package main_test

import (
	. "code.cloudfoundry.org/backup-plugin"

	"io"

	"io/ioutil"

	"code.cloudfoundry.org/backup-plugin/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("BackupPlugin", func() {
	var reader *io.PipeReader
	var writer *io.PipeWriter
	var backupPlugin BackupPlugin

	BeforeEach(func() {
		reader, writer = io.Pipe()
		backupPlugin = BackupPlugin{Output: writer}
	})

	DescribeTable("Run",
		func(command, action string) {
			cliConnection := new(fakes.FakeCliConnection)
			cliConnection.CliCommandWithoutTerminalOutputReturnsOnCall(1, []string{"...", "Status: update in progress", "Message: ", "..."}, nil)
			cliConnection.CliCommandWithoutTerminalOutputReturnsOnCall(2, []string{"...", "Status: update in progress", "Message: ", "..."}, nil)
			cliConnection.CliCommandWithoutTerminalOutputReturnsOnCall(3, []string{"...", "Status: update succeeded", "Message: some message\nwith\nnewlines", "..."}, nil)

			go func() {
				backupPlugin.Run(cliConnection, []string{command, "service-name"})
				writer.Close()
			}()
			output, _ := ioutil.ReadAll(reader)

			Expect(cliConnection.CliCommandWithoutTerminalOutputCallCount()).To(Equal(4))
			Expect(cliConnection.CliCommandWithoutTerminalOutputArgsForCall(0)).To(Equal([]string{"update-service", "service-name", "-c", `{"action": "` + action + `"}`}))

			Expect(string(output)).To(Equal("Running " + command + "...\nsome message\nwith\nnewlines\n"))
		},
		Entry("backup-service", "backup-service", "backup"),
		Entry("restore-service", "restore-service", "restore"),
		Entry("list-service-backups", "list-service-backups", "list"),
	)

	// error in running update-service
	// error in running service
	// update-service never completing successfully
	// error in parsing the output
})
