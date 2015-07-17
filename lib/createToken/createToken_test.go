package createToken_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/strajah/strajah-token-go/lib/createToken"
	"testing"
)

var _ = Describe("Token creation", func() {
	It("Should generate a token", func() {
		data := createToken.DataToBeTokenized{
			UserId: "user-123",
		}
		token := createToken.Create(&data)

		Expect(token).ToNot(BeNil())
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Token creation")
}
