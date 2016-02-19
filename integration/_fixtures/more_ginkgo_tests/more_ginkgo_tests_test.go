package more_ginkgo_tests_test

import (
	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/ginkgo/integration/_fixtures/more_ginkgo_tests"
	. "github.com/scalingdata/gomega"
)

var _ = Describe("MoreGinkgoTests", func() {
	It("should pass", func() {
		Ω(AlwaysTrue()).Should(BeTrue())
	})

	It("should always pass", func() {
		Ω(AlwaysTrue()).Should(BeTrue())
	})
})
