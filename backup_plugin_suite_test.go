package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBackupPlugin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BackupPlugin Suite")
}
