package tags_tests_test

import (
	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"

	"testing"
)

func TestTagsTests(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TagsTests Suite")
}
