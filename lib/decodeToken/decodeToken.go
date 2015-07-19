package decodeToken

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/gob"
	"bytes"
	"github.com/strajah/strajah-token-go/lib/settings"
	"fmt"
)

type TokenSet struct {
	UserId string
}

func Decode(settings settings.Settings, encodedToken string) (TokenSet, error) {
	aesBlockDecryptor, _ := aes.NewCipher([]byte(settings.CipherKey))
	aesDecryptor := cipher.NewCFBDecrypter(aesBlockDecryptor, []byte("abcdefghabcdefgh"))

	encodedDataBytes, _ := base64.StdEncoding.DecodeString(encodedToken)
	decodedDataBytes := make([]byte, len(encodedDataBytes))

	aesDecryptor.XORKeyStream(decodedDataBytes, encodedDataBytes)

	return unserialize(decodedDataBytes)
}

func unserialize (decodedDataBytes []byte) (TokenSet, error){
	decodedDataBuffer := bytes.NewBuffer(decodedDataBytes)
	decoder := gob.NewDecoder(decodedDataBuffer)

	var decodedStructure TokenSet
	decoder.Decode(&decodedStructure)
	if decodedStructure == (TokenSet{}) {
		return decodedStructure, fmt.Errorf("Unable to decode")
	}
	return decodedStructure, nil
}
