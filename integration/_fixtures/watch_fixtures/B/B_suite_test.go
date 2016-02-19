package B_test

import (
	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"

	"testing"
)

func TestB(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "B Suite")
}
