package sdk

import (
	//"encoding/json"
	"net/http"
)

// DeliveryAddRequest
type DeliveryAddRequest struct {
	AdorderId string `json:"adorderid"`
}

// DeliveryAddResponse
type DeliveryAddResponse struct {
	BaseResponse
}

// Method
func (req *DeliveryAddRequest) Method() string {
	return http.MethodPost
}

// Path
func (req *DeliveryAddRequest) Path() string {
	return "/api/v2/delivery/add"
}

// Params
func (req *DeliveryAddRequest) Params() map[string]interface{} {
	/*
		bs, err := json.Marshal(req)
		if err != nil {
			panic(err)
		}
	*/
	m := make(map[string]interface{})
	//m["verison"] = 2
	//m["adorderid"] = string(bs[:])
	return m
}

// DelievryAdd
func (client *Client) DeliveryAdd(request *DeliveryAddRequest) (response *DeliveryAddResponse, err error) {
	response = &DeliveryAddResponse{}
	err = client.Do(request, response)
	return
}
