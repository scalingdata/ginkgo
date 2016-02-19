package coverage_fixture_test

import (
	. "github.com/scalingdata/ginkgo/integration/_fixtures/coverage_fixture"

	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"
)

var _ = Describe("CoverageFixture", func() {
	It("should test A", func() {
		Ω(A()).Should(Equal("A"))
	})

	It("should test B", func() {
		Ω(B()).Should(Equal("B"))
	})

	It("should test C", func() {
		Ω(C()).Should(Equal("C"))
	})

	It("should test D", func() {
		Ω(D()).Should(Equal("D"))
	})
})
