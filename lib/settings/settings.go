package settings

import (
	"fmt"
)

type Settings struct {
	CipherKey string
}

func (customSettings Settings) Check () (error) {
	cipherKeyBytes := []byte(customSettings.CipherKey)
	cipherKeyBytesLen := len(cipherKeyBytes)

	if (cipherKeyBytesLen == 16 || cipherKeyBytesLen == 24 || cipherKeyBytesLen == 32) {
		return nil

	}

	return fmt.Errorf("cipher key is not 16 or 24 bytes long")
}
