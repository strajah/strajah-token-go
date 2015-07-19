package decodeToken

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/strajah/strajah-token-go/lib/createToken"
)

var _ = Describe("Decode tokens", func() {
	It("Generated token must be decoded back to get original data", func() {
		userId := "user-123"
		data := createToken.DataToBeTokenized{
			UserId: userId,
		}
		token := createToken.Create(&data)

		tokenSet := Decode(token)
		Expect(tokenSet.UserId).To(Equal(userId))
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Decode tokens")
}
