package tools_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/rmwiesenberg/ode/tools"
)

var _ = Describe("Language", func() {
	It("should return a verbed phrase", func() {
		testPhrase := "to the moon"

		response := tools.Verber(testPhrase)
		Expect(response).To(Equal("toes the moon"))
	})
})
