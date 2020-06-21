package functions

import (
	"fmt"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"encoding/base64"
	st "app/structs"
)
// decrypt from base64 to decrypted string
func Decrypt(key []byte, text string) string {
	textEncode, _ := base64.URLEncoding.DecodeString(text)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "-1"
	}
	if len(textEncode) < aes.BlockSize {
		return "-1"
	} 
	iv := textEncode[:aes.BlockSize]
	textEncode = textEncode[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(textEncode, textEncode)
	result := &st.RequestBodyReturn{
		Code: 200,
		Message: "Success",
		Data: fmt.Sprintf("%s", textEncode)}
	res, _ := json.Marshal(result)
	return string(res)
	
}
