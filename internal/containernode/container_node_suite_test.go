package containernode_test

import (
	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"

	"testing"
)

func TestContainernode(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Containernode Suite")
}
