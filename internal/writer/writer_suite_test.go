package writer_test

import (
	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"

	"testing"
)

func TestWriter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Writer Suite")
}
