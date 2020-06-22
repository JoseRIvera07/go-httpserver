package Handlers

import (
	"fmt"
	"net/http"
    "encoding/json"
    st"../structs"
)

// MyError contains custom code, error message, and HTTP status code.
type MyError struct {
    HTTPStatus int    `json:"-"`
    Code       int    `json:"code"`
    Message    string `json:"message"`
}

func (e *MyError) Error() string {
    return fmt.Sprintf("HTTPStatus: %v, Code: %v, Message: %q",
        e.HTTPStatus, e.Code, e.Message)
}

// WriteToResponse writes response for the error.
func (e *MyError) WriteToResponse(w http.ResponseWriter) {
    w.WriteHeader(e.HTTPStatus)
    fmt.Fprintf(w, e.ToJSON())
}

// ToJSON returns JSON string for a MyError.
func (e *MyError) ToJSON() string {
    j, err := json.Marshal(e)
    if err != nil {
        return `{"code":500,"message":json.Marshal() failed"}`
    }
    return string(j)
}

func MethodNotAllowedErr(method string) *MyError {
    return &MyError{
        HTTPStatus: http.StatusMethodNotAllowed,
        Code:       st.Config.MethodNotAllowedCode,
        Message:    fmt.Sprintf("method %q is not allowed", method),
    }
}
// JSONDecoderErr .
func InternalErrorCode(err error) *MyError {
    return &MyError{
        HTTPStatus: http.StatusInternalServerError,
        Code:       st.Config.InternalErrorCode,
        Message:    err.Error(),
    }
}

//DecryptErr
func DecryptErr(text string) *MyError {
    return &MyError{
        HTTPStatus: http.StatusInternalServerError,
        Code:       st.Config.InternalErrorCode,
        Message:    fmt.Sprintf("text %q is not correctly encryted", text),
    }
}
