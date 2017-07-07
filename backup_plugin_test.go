package main_test

import (
	. "code.cloudfoundry.org/backup-plugin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("BackupPlugin", func() {
	var backupPlugin = BackupPlugin{}

	Describe("GetMetadata", func() {
		It("Returns a PluginMetadata struct", func() {
			metadata := backupPlugin.GetMetadata()

			Expect(metadata.Name).To(Equal("BackupPlugin"))
		})
	})
})
