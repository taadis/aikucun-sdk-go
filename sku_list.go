package aikucun

import (
	"net/http"
)

// SkuListRequest
type SkuListRequest struct {
	ActivityId string `json:"activityid"`
	Page       int    `json:"page"`
	PageSize   int    `json:"pagesize"`
}

// SkuListResponse
type SkuListResponse struct {
	BaseResponse
	SkuListResponseData struct {
		TotalRecord int `json:"totalrecord"`
		PageSize    int `json:"pagesize"`
		TotalPage   int `json:"totalpage"`
		Page        int `json:"page"`
		List        []struct {
			Id        string `json:"id"`
			BarCode   string `json:"barcode"`
			KuanHao   string `json:"kuanhao"`
			BrandId   string `json:"brandid"`
			BrandName string `json:"brandname"`
			Name      string `json:"name"`
			Size      string `json:"size"`
			Color     string `json:"color"`
		}
	}
}

// Method
func (req *SkuListRequest) Method() string {
	return http.MethodGet
}

// Path
func (req *SkuListRequest) Path() string {
	return "/api/v2/sku/list"
}

// Params
func (req *SkuListRequest) Params() map[string]interface{} {
	m := make(map[string]interface{})
	m["activityid"] = req.ActivityId
	m["page"] = req.Page
	m["pagesize"] = req.PageSize
	return m
}

// NewSkuListRequest
func NewSkuListRequest() *SkuListRequest {
	return new(SkuListRequest)
}

// SkuList
func (client *Client) SkuList(request *SkuListRequest) (response *SkuListResponse, err error) {
	response = &SkuListResponse{}
	err = client.Do(request, response)
	return
}
