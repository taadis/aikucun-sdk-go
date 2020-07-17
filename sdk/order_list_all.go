package sdk

import (
	"net/http"
)

// OrderListAllRequest
type OrderListAllRequest struct {
	ActivityId  string `json:"activityid"`
	Page        int    `json:"page"`
	PageSize    int    `json:"pagesize"`
	WithWaybill int    `json:"withwaybill"`
}

// OrderListAllResponse
type OrderListAllResponse struct {
}

// Method
func (req *OrderListAllRequest) Method() string {
	return http.MethodGet
}

// Path
func (req *OrderListAllRequest) Path() string {
	return "/api/v2/order/listall?version=2"
}

// Params
func (req *OrderListAllRequest) Params() map[string]interface{} {
	m := make(map[string]interface{})
	m["activityid"] = req.ActivityId
	m["page"] = req.Page
	m["pagesize"] = req.PageSize
	m["withwaybill"] = req.WithWaybill
	return m
}

// OrderListAll
func (client *Client) OrderListAll(request *OrderListAllRequest) (response *OrderListAllResponse, err error) {
	response = &OrderListAllResponse{}
	// TODO: err = client.Do(request, response)
	return
}
