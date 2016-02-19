package passing_before_suite_test

import (
	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"
)

var _ = Describe("PassingSuiteSetup", func() {
	It("should pass", func() {
		Ω(a).Should(Equal("ran before suite"))
		Ω(b).Should(BeEmpty())
	})

	It("should pass", func() {
		Ω(a).Should(Equal("ran before suite"))
		Ω(b).Should(BeEmpty())
	})

	It("should pass", func() {
		Ω(a).Should(Equal("ran before suite"))
		Ω(b).Should(BeEmpty())
	})

	It("should pass", func() {
		Ω(a).Should(Equal("ran before suite"))
		Ω(b).Should(BeEmpty())
	})
})
