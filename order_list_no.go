package aikucun

import (
	"net/http"
)

// OrderListNoRequest
type OrderListNoRequest struct {
	ActivityId    string `json:"activityid"`
	Page          int    `json:"page"`
	PageSize      int    `json:"pagesize"`
	FilterDeliver int    `json:"filterDeliver"`
}

// OrderListNoResponse
type OrderListNoResponse struct {
	BaseResponse
	Data OrderListNoData `json:"data"`
}

// OrderListNoData
type OrderListNoData struct {
	TotalRecord int               `json:"totalrecord"`
	TotalPage   int               `json:"totalpage"`
	Page        int               `json:"page"`
	PageSize    int               `json:"pagesize"`
	Source      string            `json:"source"`
	ActivityId  string            `json:"activityid"`
	List        []OrderListNoItem `json:"list"`
}

// OrderListNoItem
type OrderListNoItem struct {
	AdorderId                 string `json:"adorderid"`
	RealSettlementpriceAmount int    `json:"realSettlementpriceAmount"`
	Count                     int    `json:"count"`
	OrderStatus               int    `json:"orderstatus"`
	LogisticsStatus           int    `json:"logisticsstatus"`
	OrderTime                 string `json:"ordertime"`
	PayTime                   string `json:"paytime"`
}

// Method
func (req *OrderListNoRequest) Method() string {
	return http.MethodGet
}

// Path
func (req *OrderListNoRequest) Path() string {
	return "/api/v2/order/listno"
}

// Params
func (req *OrderListNoRequest) Params() map[string]interface{} {
	m := make(map[string]interface{})
	m["activityid"] = req.ActivityId
	m["page"] = req.Page
	m["pagesize"] = req.PageSize
	m["filterDeliver"] = req.FilterDeliver
	return m
}

// NewOrderListNoRequest
func NewOrderListNoRequest() *OrderListNoRequest {
	return new(OrderListNoRequest)
}

// OrderListNo
func (client *Client) OrderListNo(request *OrderListNoRequest) (response *OrderListNoResponse, err error) {
	response = &OrderListNoResponse{}
	err = client.Do(request, response)
	return
}
