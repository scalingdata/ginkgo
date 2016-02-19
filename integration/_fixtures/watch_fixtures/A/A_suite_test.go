package A_test

import (
	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"

	"testing"
)

func TestA(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "A Suite")
}
