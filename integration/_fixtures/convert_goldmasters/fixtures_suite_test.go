package tmp

import (
	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"

	"testing"
)

func TestTmp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tmp Suite")
}
