package structs

type Configuration struct {
	Key string
	HttpMethod string
	MethodNotAllowedCode int
	JSONDecoderCode int
	SuccessCode int
}

var Config = Configuration{}