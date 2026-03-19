package encoding

import "encoding/base64"

func Base64Decode(encryptedStr string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(encryptedStr)
}

func MustBase64Decode(encryptedStr string) []byte {
	decoded, err := Base64Decode(encryptedStr)
	if err != nil {
		panic(err)
	}
	return decoded
}
