package eventually_failing_test

import (
	"fmt"
	"io/ioutil"
	"time"

	. "github.com/scalingdata/ginkgo"
	. "github.com/scalingdata/gomega"
)

var _ = Describe("EventuallyFailing", func() {
	It("should fail on the third try", func() {
		time.Sleep(time.Second)
		files, err := ioutil.ReadDir(".")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(len(files)).Should(BeNumerically("<", 5))
		ioutil.WriteFile(fmt.Sprintf("./%d", len(files)), []byte("foo"), 0777)
	})
})
