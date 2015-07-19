package createToken

import (
	"encoding/gob"
	"encoding/base64"
	"crypto/cipher"
	"crypto/aes"
	"bytes"
	"github.com/strajah/strajah-token-go/lib/settings"
)

type DataToBeTokenized struct {
	UserId string
}

func Create (settings settings.Settings, data *DataToBeTokenized) (string) {
	aesCipherBlock, _ := aes.NewCipher([]byte(settings.CipherKey))
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
