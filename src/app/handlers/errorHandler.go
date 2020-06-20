package handlers

import (
	"fmt"
	"net/http"
    "encoding/json"
    st"app/structs"
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
        return `{"code":500,"message":"ScrapError.JSONStr: json.Marshal() failed"}`
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
func JSONDecoderErr(err error) *MyError {
    return &MyError{
        HTTPStatus: http.StatusInternalServerError,
        Code:       st.Config.JSONDecoderCode,
        Message:    err.Error(),
    }
}
