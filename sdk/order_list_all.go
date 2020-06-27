package sdk

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

// OrderListAll
func (client *Client) OrderListAll(request *OrderListAllRequest) (response *OrderListAllResponse, err error) {
	response = &OrderListAllResponse{}
	// TODO: err = client.Do(request, response)
	return
}
