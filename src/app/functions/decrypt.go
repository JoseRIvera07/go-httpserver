package functions

import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"encoding/base64"
	st "../structs"
)
// Decrypt from base64 encrypted text 
func Decrypt(key []byte, text string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(text)

	// Generate a new aes cipher with the key
	// key parameter must be 16, 24 or 32, corresponding to the AES-128, AES-192 or AES-256 algorithms, respectively
	block, err := aes.NewCipher(key)
	if err != nil {
		return "-1"
	}
	if len(ciphertext) < aes.BlockSize {
		return "-1"
	} 
	// Apply decryption
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	// Create the json response
	result := &st.RequestBodyReturn{
		Code: 200,
		Message: "Success",
		Data: fmt.Sprintf("%s", ciphertext)}
	res, _ := json.Marshal(result)
	return string(res)
	
}
