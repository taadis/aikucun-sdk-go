package sdk

// DeliveryAddRequest
type DeliveryAddRequest struct {
	AdorderId string `json:"adorderid"`
}

// DeliveryAddResponse
type DeliveryAddResponse struct {
}

// DelievryAdd
func (this *Client) DeliveryAdd(request *DeliveryAddRequest) (response *DeliveryAddResponse, err error) {
	response = &DeliveryAddResponse{}
	// TODO:
	return
}
