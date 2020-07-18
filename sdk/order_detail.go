package sdk

import (
	"net/http"
)

// OrderDetailRequest
type OrderDetailRequest struct {
	AdorderId   string `json:"adorderid"`
	WithWaybill int    `json:"withwaybill"`
}

// Method
func (req *OrderDetailRequest) Method() string {
	return http.MethodGet
}

// Path
func (req *OrderDetailRequest) Path() string {
	return "/api/v2/order/detail"
}

// Params
func (req *OrderDetailRequest) Params() map[string]interface{} {
	m := make(map[string]interface{})
	m["adorderid"] = req.AdorderId
	m["withwaybill"] = req.WithWaybill
	return m
}

// OrderDetailResponse
type OrderDetailResponse struct {
	BaseResponse
	Data OrderDetailData `json:"data"`
}

// OrderDetailData
type OrderDetailData struct {
	AdorderId                 string               `json:"adorderid"`
	Amount                    int                  `json:"amount"`
	RealSettlementpriceAmount string               `json:"realSettlementpriceAmount"`
	Count                     int                  `json:"count"`
	OrderStatus               int                  `json:"orderstatus"`
	LogisticsStatus           int                  `json:"logisticsstatus"`
	OrderTime                 string               `json:"ordertime"`
	PayTime                   string               `json:"paytime"`
	Delivery                  OrderDetailDelivery  `json:"delivery"`
	List                      []OrderDetailProduct `json:"list"`
}

// OrderDetailDelivery
type OrderDetailDelivery struct {
	SkuId                string `json:"skuid"`
	DeliverNo            string `json:"deliverNo"`
	LogisticsCompany     string `json:"logisticsCompany"`
	InsuranceValue       int    `json:"insuranceValue"`
	BackSignBill         string `json:"backSignBill"`
	CodValue             int    `json:"codValue"`
	Province             string `json:"province"`
	City                 string `json:"city"`
	County               string `json:"county"`
	ReceiverArea         string `json:"receiverArea"`
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

// OrderDetailProduct
type OrderDetailProduct struct {
	SkuId               string `json:"skuid"`
	BarCode             string `json:"barcode"`
	BrandId             string `json:"brandid"`
	BrandName           string `json:"brandname"`
	Name                string `json:"name"`
	Size                string `json:"size"`
	Color               string `json:"color"`
	KuanHao             string `json:"kuanhao"`
	Num                 int    `json:"num"`
	RealSettlementPrice string `json:"realsettlementprice"`
}

// OrderDetail
func (client *Client) OrderDetail(request *OrderDetailRequest) (response *OrderDetailResponse, err error) {
	response = &OrderDetailResponse{}
	err = client.Do(request, response)
	return
}
