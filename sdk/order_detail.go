package sdk

// OrderDetailRequest
type OrderDetailRequest struct {
	AdorderId   string `json:"adorderid"`
	WithWaybill int    `json:"withwaybill"`
}

// OrderDetailResponse
type OrderDetailResponse struct {
}

// OrderDetail
func (client *Client) OrderDetail(request *OrderDetailRequest) (response *OrderDetailResponse, err error) {
	response = &OrderDetailResponse{}
	// TODO:
	return
}
