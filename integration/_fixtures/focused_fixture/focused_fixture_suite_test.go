package focused_fixture_test

import (
	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"

	"testing"
)

func TestFocused_fixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Focused_fixture Suite")
}
