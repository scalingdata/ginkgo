package C_test

import (
	. "github.com/scalingdata/ginkgo/integration/_fixtures/watch_fixtures/C"

	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"
)

var _ = Describe("C", func() {
	It("should do it", func() {
		Ω(DoIt()).Should(Equal("done!"))
	})
})
