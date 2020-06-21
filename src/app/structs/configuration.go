package structs

type Configuration struct {
	Ip string
	Port string
	Key string
	EncryptRoute string
	DecryptRoute string
	HttpMethod string
	MethodNotAllowedCode int
	InternalErrorCode int
	SuccessCode int
}

var Config = Configuration{}