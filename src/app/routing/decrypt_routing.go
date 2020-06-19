package routing

import (
	"fmt"
	"net/http"
	"encoding/json"
	fn "app/functions"
)

type Response_Decrypt_Routing struct {
	Data string
}

func Decrypt_Routing(w http.ResponseWriter, request *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(request.Body)
	var t Response_Decrypt_Routing

	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}
	key := []byte("example key 1234")
	text := fn.Decrypt(key, t.Data)
	fmt.Fprint(w, text)
	
}