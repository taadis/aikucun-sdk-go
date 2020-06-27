package sdk

// OrderListNoRequest
type OrderListNoRequest struct {
	ActivityId    string `json:"activityid"`
	Page          int    `json:"page"`
	PageSize      int    `json:"pagesize"`
	FilterDeliver int    `json:"filterDeliver"`
}

// OrderListNoResponse
type OrderListNoResponse struct {
}

// OrderListNo
func (client *Client) OrderListNo(request *OrderListNoRequest) (response *OrderListNoResponse, err error) {
	response = &OrderListNoResponse{}
	// TODO:
	return
}
