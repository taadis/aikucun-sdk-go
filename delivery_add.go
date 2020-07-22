package aikucun

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
	Data DeliveryAddData `json:"data"`
}

// DeliveryAddData
type DeliveryAddData struct {
	DeliveryNo           string `json:"deliverNo"`
	LogisticsCompany     int    `json:"logisticsCompany"`
	InsuranceValue       int    `json:"insuranceValue"`
	BackSignBill         string `json:"backSignBill"`
	CodValue             int    `json:"codValue"`
	Province             string `json:"province"`
	City                 string `json:"city"`
	County               string `json:"county"`
	ReciverArea          string `json:"receiverArea"`
	Receiver             string `json:"receiver"`
	ReceiverTel          string `json:"receiverTel"`
	ReceiverAddress      string `json:"receiverAddress"`
	Sender               string `json:"sender"`
	SendTel              string `json:"sendTel"`
	SendAddress          string `json:"sendAddress"`
	Remark               string `json:"remark"`
	SectionCode          string `json:"sectionCode"`
	SourceSortCenterName string `json:"sourceSortCenterName"`
	TargetSortCenterName string `json:"targetSortCenterName"`
	DestinationCity      string `json:"destinationCity"`
	LogisticsAreaCode    string `json:"logisticsAreaCode"`
	Col1                 string `json:"col1"`
	Col2                 string `json:"col2"`
	Col3                 string `json:"col3"`
	Col4                 string `json:"col4"`
	Col5                 string `json:"col5"`
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
	m := make(map[string]interface{})
	m["version"] = 2
	return m
}

// DelievryAdd
func (client *Client) DeliveryAdd(request *DeliveryAddRequest) (response *DeliveryAddResponse, err error) {
	response = &DeliveryAddResponse{}
	err = client.Do(request, response)
	return
}
