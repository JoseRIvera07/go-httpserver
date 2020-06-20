package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	fn "app/functions"
	st "app/structs"
)

func EncryptHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")  
	w.Header().Set("Content-Type", "application/json")
	if request.Method != st.Config.HttpMethod {
		MethodNotAllowedErr(request.Method).WriteToResponse(w)
		return
	}
	var t st.RequestBody
	decoder :=  json.NewDecoder(request.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&t)
	if  err != nil {
		InternalErrorCode(err).WriteToResponse(w)
		return
	}
	key := []byte(st.Config.Key)
	resultEncrypText := fn.Encrypt(key, t.Data)
	fmt.Fprint(w, resultEncrypText)
}