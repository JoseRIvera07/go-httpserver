package Handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	st "../structs"
    fn "../functions"
)
// Function to handle Decrypt functions
func DecryptHandler(w http.ResponseWriter, request *http.Request) {
	// Set Header to send json responde to the cient.
	w.Header().Set("Content-Type", "application/json")
	// Validation to accept only the request comming from [POST] method
	// It could have benn done using .Method("POST"), but I used this way to respond using personalized json.
	if request.Method != st.Config.HttpMethod {
		MethodNotAllowedErr(request.Method).WriteToResponse(w)
		return
	}
	var t st.RequestBody
	// Decode the received json
	decoder :=  json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&t)
	// Validate json, if the msg received is not a json or does not have the "data" key.
	if  err != nil {
		InternalErrorCode(err).WriteToResponse(w)
		return
	}
	key := []byte(st.Config.Key)
	// Decrypt received text
	resultDecryptText := fn.Decrypt(key, t.Data)
	// Validate if decryption was succesful or not
	// If yes, Decrypt functions returns a json with decrypted text and http successful code
	// If not, returns -1
	if resultDecryptText == "-1"{
		DecryptErr(t.Data).WriteToResponse(w)
		return
	}
	fmt.Fprint(w, resultDecryptText)
}