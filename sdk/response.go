package sdk

// IResponse
type IResponse interface {
}

// BaseResponse
type BaseResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}
