package functions

import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"encoding/base64"
	rq "app/structs"
)
// decrypt from base64 to decrypted string
func Decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	if len(ciphertext) < aes.BlockSize {
		panic("text too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	result := &rq.RequestBodyReturn{
		Code: 200,
		Message: "Success",
		Data: fmt.Sprintf("%s", ciphertext)}
	res, _ := json.Marshal(result)
	return string(res)
}