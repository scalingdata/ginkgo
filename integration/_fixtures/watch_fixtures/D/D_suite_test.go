package D_test

import (
	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"

	"testing"
)

func TestD(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "D Suite")
}
