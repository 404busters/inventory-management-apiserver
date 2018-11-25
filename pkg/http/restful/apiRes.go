package restful

type ErrorRes struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Params  map[string]interface{} `json:"params,omitempty"`
}

type ApiRes struct {
	Data interface{} `json:"data"`
}
