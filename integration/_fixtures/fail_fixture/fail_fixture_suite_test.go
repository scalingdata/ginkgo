package fail_fixture_test

import (
	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"

	"testing"
)

func TestFail_fixture(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fail_fixture Suite")
}
