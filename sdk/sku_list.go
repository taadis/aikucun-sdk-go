package sdk

import (
	"net/http"
	"strconv"
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
func (req *SkuListRequest) Params() map[string]string {
	m := make(map[string]string)
	m["activityid"] = req.ActivityId
	m["page"] = strconv.Itoa(req.Page)
	m["pagesize"] = strconv.Itoa(req.PageSize)
	return m
}

// SkuList
func (client *Client) SkuList(request *SkuListRequest) (response *SkuListResponse, err error) {
	response = &SkuListResponse{}
	err = client.Do(request, response)
	return
}
