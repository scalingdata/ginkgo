package failing_ginkgo_tests_test

import (
	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/ginkgo/integration/_fixtures/failing_ginkgo_tests"
	. "github.com/scalingdata/gomega"
)

var _ = Describe("FailingGinkgoTests", func() {
	It("should fail", func() {
		Ω(AlwaysFalse()).Should(BeTrue())
	})

	It("should pass", func() {
		Ω(AlwaysFalse()).Should(BeFalse())
	})
})
