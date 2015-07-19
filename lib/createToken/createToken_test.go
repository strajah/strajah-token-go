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
		token, _ := Create(settings, &data)

		Expect(token).ToNot(BeNil())
	})

	It("Should not allow token creation if cipher key is different than 16, 24, or 32 bytes", func() {
		data := DataToBeTokenized{
			UserId: "User-12345",
		}

		settings := settings.Settings{CipherKey:"12345678"} //length = 8
		token, err := Create(settings, &data)
		Expect(token).To(Equal(""))
		Expect(err).ToNot(BeNil())

		settings.CipherKey = "12345678123456781234" //length = 20
		token, err = Create(settings, &data)
		Expect(token).To(Equal(""))
		Expect(err).ToNot(BeNil())

		settings.CipherKey = "1234567812345678123412345678123456781234" //length = 40
		token, err = Create(settings, &data)
		Expect(token).To(Equal(""))
		Expect(err).ToNot(BeNil())
	})

	It("Should allow token creation for cipher keys with byte length 16, 24 or 32", func() {
		data := DataToBeTokenized{
			UserId: "User-12345",
		}

		settings := settings.Settings{CipherKey:"1234567812345678"} //length = 16
		_, err := Create(settings, &data)
		Expect(err).To(BeNil())

		settings.CipherKey = "123456781234567812345678" //length = 24
		_, err = Create(settings, &data)
		Expect(err).To(BeNil())

		settings.CipherKey = "12341234123412341234123412341234" //length = 32
		_, err = Create(settings, &data)
		Expect(err).To(BeNil())
	})
})

func Test(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Token creation")
}
