package functions

import (
	"io"
	"crypto/aes"
	"crypto/rand"
	"crypto/cipher"
	"encoding/json"
	"encoding/base64"
	rq "../structs"
)
// encrypt string to base64 crypto using AES
func Encrypt(key []byte, text string) string {
	plaintext := []byte(text)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "-1"
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	// Stream that has encrypts with cipher feedback mode,
	// using the given Block. The iv must be the same length as the Block's block
	// size.
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	// Create the json response
	result := &rq.RequestBodyReturn{
		Code: 200,
		Message: "Success",
		Data: base64.URLEncoding.EncodeToString(ciphertext)}
	res, _ := json.Marshal(result)

	return string(res)
}
