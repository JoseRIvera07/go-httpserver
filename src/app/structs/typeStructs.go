package structs

type RequestBody struct {
	Data string `json:"data"`
}

type RequestBodyReturn struct {
	Code       int    `json:"code"`
    Message    string `json:"message"`
	Data string `json:"data"`
}