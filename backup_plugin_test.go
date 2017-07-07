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
		func(cliArgs, expectedUpdateServiceArgs []string) {
			cliConnection := new(fakes.FakeCliConnection)
			cliConnection.CliCommandWithoutTerminalOutputReturnsOnCall(1, []string{"...", "Status: update in progress", "Message: ", "..."}, nil)
			cliConnection.CliCommandWithoutTerminalOutputReturnsOnCall(2, []string{"...", "Status: update in progress", "Message: ", "..."}, nil)
			cliConnection.CliCommandWithoutTerminalOutputReturnsOnCall(3, []string{"...", "Status: update succeeded", "Message: some message\nwith\nnewlines", "..."}, nil)

			go func() {
				backupPlugin.Run(cliConnection, cliArgs)
				writer.Close()
			}()
			output, _ := ioutil.ReadAll(reader)

			Expect(cliConnection.CliCommandWithoutTerminalOutputCallCount()).To(Equal(4))
			Expect(cliConnection.CliCommandWithoutTerminalOutputArgsForCall(0)).To(Equal(expectedUpdateServiceArgs))

			Expect(string(output)).To(Equal("Running " + cliArgs[0] + "...\nsome message\nwith\nnewlines\n"))
		},
		Entry("backup-service", []string{"backup-service", "service-name"}, []string{"update-service", "service-name", "-c", `{"action": "backup"}`}),
		Entry("restore-service", []string{"restore-service", "service-name", "backup-id"}, []string{"update-service", "service-name", "-c", `{"action": "restore", "backup_id": "backup-id"}`}),
		Entry("list-service-backups", []string{"list-service-backups", "service-name"}, []string{"update-service", "service-name", "-c", `{"action": "list"}`}),
	)

	// error in running update-service
	// error in running service
	// update-service never completing successfully
	// error in parsing the output
})
