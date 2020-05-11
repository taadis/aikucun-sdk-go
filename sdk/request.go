package sdk

// BaseRequest
type BaseRequest struct {
	AppId      string `json:"appid"`
	NonceStr   string `json:"noncestr"`
	Timestamp  int64  `json:"timestamp"`
	Erp        string `json:"erp"`
	ErpVerison string `json:"erpversion"`
	Sign       string `json:"sign"`
}
