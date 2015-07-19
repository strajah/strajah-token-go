package createToken

import (
	"encoding/gob"
	"encoding/base64"
	"crypto/cipher"
	"crypto/aes"
	"bytes"
)

type DataToBeTokenized struct {
	UserId string
}

func Create (data *DataToBeTokenized) (string) {
	aesCipherBlock, _ := aes.NewCipher([]byte("0123456789ABCDEF"))
	aesEncrypter := cipher.NewCFBEncrypter(aesCipherBlock, []byte("abcdefghabcdefgh"))

	serializedData := serialize(data)
	encryptedData := make([]byte, len(serializedData.Bytes()))

	aesEncrypter.XORKeyStream(encryptedData, serializedData.Bytes())
	return base64.StdEncoding.EncodeToString(encryptedData)
}

func serialize (dataToBeSerialized *DataToBeTokenized) (buffer bytes.Buffer){
	encoder := gob.NewEncoder(&buffer)
	encoder.Encode(dataToBeSerialized)
	return
}
