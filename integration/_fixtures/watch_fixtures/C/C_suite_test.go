package C_test

import (
	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"

	"testing"
)

func TestC(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "C Suite")
}
