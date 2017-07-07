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
			cliConnection.CliCommandReturnsOnCall(1, []string{"...", "Status: update in progress", "Message: ", "..."}, nil)
			cliConnection.CliCommandReturnsOnCall(2, []string{"...", "Status: update in progress", "Message: ", "..."}, nil)
			cliConnection.CliCommandReturnsOnCall(3, []string{"...", "Status: update succeeded", "Message: some message\nwith\nnewlines", "..."}, nil)

			go func() {
				backupPlugin.Run(cliConnection, []string{command, "service-name"})
				writer.Close()
			}()
			output, _ := ioutil.ReadAll(reader)

			Expect(cliConnection.CliCommandCallCount()).To(Equal(4))
			Expect(cliConnection.CliCommandArgsForCall(0)).To(Equal([]string{"update-service", "service-name", "-c", `{"backup-action": "` + action + `"}`}))

			Expect(string(output)).To(Equal("some message\nwith\nnewlines\n"))
		},
		Entry("backup-service", "backup-service", "backup"),
		Entry("restore-service", "restore-service", "restore"),
		Entry("list-service-backups", "list-service-backups", "list"),
	)
})
