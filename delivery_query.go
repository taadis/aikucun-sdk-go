package aikucun

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// DeliveryQueryRequest
type DeliveryQueryRequest struct {
	AdorderId string `json:"adorderid"`
}

// DeliveryQueryResponse
type DeliveryQueryResponse struct {
	BaseResponse
	Data DeliveryQueryData `json:"data"`
}

// DeliveryQueryData
type DeliveryQueryData struct {
	DeliverNos           []string `json:"deliverNos"`
	LogisticsCompany     int      `json:"logisticsCompany"`
	InsuranceValue       int      `json:"insuranceValue"`
	BackSignBill         string   `json:"backSignBill"`
	CodValue             int      `json:"codValue"`
	Province             string   `json:"province"`
	City                 string   `json:"city"`
	County               string   `json:"county"`
	ReceiverArea         string   `json:"receiverArea"`
	Receiver             string   `json:"receiver"`
	ReceiverTel          string   `json:"receiverTel"`
	ReceiverAddress      string   `json:"receiverAddress"`
	Sender               string   `json:"sender"`
	SendTel              string   `json:"sendTel"`
	SendAddress          string   `json:"sendAddress"`
	Remark               string   `json:"remark"`
	SectionCode          string   `json:"sectionCode"`
	SourceSortCenterName string   `json:"sourceSortCenterName"`
	TargetSortCenterName string   `json:"targetSortCenterName"`
	DestinationCity      string   `json:"destinationCity"`
	LogisticsAreaCode    string   `json:"logisticsAreaCode"`
	Col1                 string   `json:"col1"`
	Col2                 string   `json:"col2"`
	Col3                 string   `json:"col3"`
	Col4                 string   `json:"col4"`
	Col5                 string   `json:"col5"`
}

// Method
func (req *DeliveryQueryRequest) Method() string {
	return http.MethodGet
}

// Path
func (req *DeliveryQueryRequest) Path() string {
	return "/api/v2/delivery/query"
}

// Params
func (req *DeliveryQueryRequest) Params() map[string]interface{} {
	m := make(map[string]interface{})
	m["adorderid"] = req.AdorderId
	m["version"] = 2
	return m
}

// NewDeliveryQueryRequest
func NewDeliveryQueryRequest() *DeliveryQueryRequest {
	return new(DeliveryQueryRequest)
}

// String
func (resp *DeliveryQueryResponse) String() string {
	buf := bytes.NewBuffer(nil)
	err := json.NewEncoder(buf).Encode(resp)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

// DeliveryQuery
func (client *Client) DeliveryQuery(request *DeliveryQueryRequest) (response *DeliveryQueryResponse, err error) {
	response = &DeliveryQueryResponse{}
	err = client.Do(request, response)
	return
}
