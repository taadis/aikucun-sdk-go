package sdk

// DeliveryQueryRequest
type DeliveryQueryRequest struct {
	AdorderId string `json:"adorderid"`
}

// DeliveryQueryResponse
type DeliveryQueryResponse struct {
}

// DeliveryQuery
func (client *Client) DeliveryQuery(request *DeliveryQueryRequest) (response *DeliveryQueryResponse, err error) {
	response = &DeliveryQueryResponse{}
	// TODO:
	return
}
