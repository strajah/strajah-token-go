package decodeToken

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/gob"
	"bytes"
)

type TokenSet struct {
	UserId string
}

func Decode(encodedToken string) (TokenSet) {
	aesBlockDecryptor, _ := aes.NewCipher([]byte("0123456789ABCDEF"))
	aesDecryptor := cipher.NewCFBDecrypter(aesBlockDecryptor, []byte("abcdefghabcdefgh"))

	encodedDataBytes, _ := base64.StdEncoding.DecodeString(encodedToken)
	decodedDataBytes := make([]byte, len(encodedDataBytes))

	aesDecryptor.XORKeyStream(decodedDataBytes, encodedDataBytes)

	return unserialize(decodedDataBytes)
}

func unserialize (decodedDataBytes []byte) (TokenSet){
	decodedDataBuffer := bytes.NewBuffer(decodedDataBytes)
	decoder := gob.NewDecoder(decodedDataBuffer)
	var decodedStructure TokenSet
	decoder.Decode(&decodedStructure)
	return decodedStructure
}
