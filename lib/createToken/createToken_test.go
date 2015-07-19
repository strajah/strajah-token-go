package createToken

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("Token creation", func() {
	It("Should generate a token", func() {
		data := DataToBeTokenized{
			UserId: "user-123",
		}
		token := Create(&data)

		Expect(token).ToNot(BeNil())
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Token creation")
}
