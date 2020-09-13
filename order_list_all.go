package aikucun

import (
	"bytes"
	"encoding/json"
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
	BaseResponse
	// ...
}

// Method
func (req *OrderListAllRequest) Method() string {
	return http.MethodGet
}

// Path
func (req *OrderListAllRequest) Path() string {
	return "/api/v2/order/listall"
}

// Params
func (req *OrderListAllRequest) Params() map[string]interface{} {
	m := make(map[string]interface{})
	m["version"] = 2
	m["activityid"] = req.ActivityId
	m["page"] = req.Page
	m["pagesize"] = req.PageSize
	m["withwaybill"] = req.WithWaybill
	return m
}

// NewOrderListAllRequest
func NewOrderListAllRequest() *OrderListAllRequest {
	return new(OrderListAllRequest)
}

// String
func (resp *OrderListAllResponse) String() string {
	buf := bytes.NewBuffer(nil)
	err := json.NewEncoder(buf).Encode(resp)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

// OrderListAll
func (client *Client) OrderListAll(request *OrderListAllRequest) (response *OrderListAllResponse, err error) {
	response = &OrderListAllResponse{}
	err = client.Do(request, response)
	return
}
