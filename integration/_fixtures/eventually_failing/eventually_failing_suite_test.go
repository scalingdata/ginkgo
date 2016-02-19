package eventually_failing_test

import (
	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"

	"testing"
)

func TestEventuallyFailing(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "EventuallyFailing Suite")
}
