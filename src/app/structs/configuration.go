package structs

type Configuration struct {
	Key string
	HttpMethod string
	MethodNotAllowedCode int
	InternalErrorCode int
	SuccessCode int
}

var Config = Configuration{}