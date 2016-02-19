package subpackage

import (
	. "github.com/scalingdata/ginkgo"
)

var _ = Describe("Testing with Ginkgo", func() {
	It("nested sub packages", func() {
		GinkgoT().Fail(true)
	})
})
