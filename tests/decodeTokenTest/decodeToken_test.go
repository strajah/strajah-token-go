package decodeTokenTest

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/strajah/strajah-token-go/lib/createToken"
	"github.com/strajah/strajah-token-go/lib/decodeToken"
	"github.com/strajah/strajah-token-go/lib/settings"
)

var _ = Describe("Decode tokens", func() {
	It("Generated token must be decoded back to get original data", func() {
		userId := "user-123"

		settings := settings.Settings{
			CipherKey: "0123456789ABCDEF",
		}

		data := createToken.DataToBeTokenized{
			UserId: userId,
		}

		token, _ := createToken.Create(settings, &data)

		tokenSet, _ := decodeToken.Decode(settings, token)
		Expect(tokenSet.UserId).To(Equal(userId))
	})

	It("Should return an error when trying to decode with invalid cipher key", func() {
		data := createToken.DataToBeTokenized{
			UserId: "userId",
		}

		invalidSettings := settings.Settings{
			CipherKey: "4321456789ABCDEF",
		}

		validSettings := settings.Settings{
			CipherKey: "0123456789ABCDEF",
		}

		token, _ := createToken.Create(validSettings, &data)

		tokenSet, err := decodeToken.Decode(invalidSettings, token)
		Expect(tokenSet).To(Equal(decodeToken.TokenSet{})) //expect it to be an empty structure
		Expect(err).ToNot(BeNil())
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Decode tokens")
}
