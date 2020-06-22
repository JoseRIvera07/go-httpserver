package Handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	fn "../functions"
	st "../structs"
)
// Function to handle Encrypt functions
func EncryptHandler(w http.ResponseWriter, request *http.Request) { 
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
	// Encrypt received text
	resultEncrypText := fn.Encrypt(key, t.Data)
	fmt.Fprint(w, resultEncrypText)
}