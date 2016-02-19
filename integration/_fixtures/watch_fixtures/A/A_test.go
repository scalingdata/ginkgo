package A_test

import (
	. "github.com/scalingdata/ginkgo/integration/_fixtures/watch_fixtures/A"

	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"
)

var _ = Describe("A", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
