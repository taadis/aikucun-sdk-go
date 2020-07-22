package aikucun

import (
	"net/http"
)

// InvoiceUploadRequest
type InvoiceUploadRequest struct {
	ActivityId string               `json:"activityid"`
	Size       int                  `json:"size"`
	List       []InvoiceUploadOrder `json:"list"`
}

// InvoiceUploadOrder
type InvoiceUploadOrder struct {
	AdorderId        string                 `json:"adorderid"`
	ErpOrderId       string                 `json:"erporderid"`
	DeliverNo        string                 `json:"deliverNo"`
	Status           int                    `json:"status"`
	Size             int                    `json:"size"`
	List             []InvoiceUploadProduct `json:"list"`
	LogisticsCompany string                 `json:"logisticsCompany"`
}

// InvoiceUploadProduct
type InvoiceUploadProduct struct {
	PinPai  string `json:"pinpai"`
	BarCode string `json:"barcode"`
	KuanHao string `json:"kuanhao"`
	Num     int    `json:"num"`
	RealNum int    `json:"realnum"`
	LackNum int    `json:"lacknum"`
}

// InvoiceUploadResponse
type InvoiceUploadResponse struct {
	BaseResponse
}

// Method
func (req *InvoiceUploadRequest) Method() string {
	return http.MethodPost
}

// Path
func (req *InvoiceUploadRequest) Path() string {
	return "/api/v2/invoice/upload"
}

// Params
func (req *InvoiceUploadRequest) Params() map[string]interface{} {
	m := make(map[string]interface{})
	m["version"] = 2
	return m
}

// InvoiceUpload
func (client *Client) InvoiceUpload(request *InvoiceUploadRequest) (response *InvoiceUploadResponse, err error) {
	response = &InvoiceUploadResponse{}
	err = client.Do(request, response)
	return
}
