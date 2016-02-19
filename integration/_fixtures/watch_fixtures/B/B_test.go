package B_test

import (
	. "github.com/scalingdata/ginkgo/integration/_fixtures/watch_fixtures/B"

	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"
)

var _ = Describe("B", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
