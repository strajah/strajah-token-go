package createToken

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
	"github.com/strajah/strajah-token-go/lib/settings"
)

var _ = Describe("Token creation", func() {
	It("Should generate a token", func() {
		settings := settings.Settings{
			CipherKey: "0123456789ABCDEF",
		}

		data := DataToBeTokenized{
			UserId: "user-123",
		}
		token := Create(settings, &data)

		Expect(token).ToNot(BeNil())
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Token creation")
}
