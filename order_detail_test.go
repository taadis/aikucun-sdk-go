package aikucun

import (
	"testing"
)

//
func TestOrderDetailResponseString(t *testing.T) {
	resp := &OrderDetailResponse{}
	s := resp.String()
	t.Log(s)
}

// TestOrderDetail
func TestOrderDetail(t *testing.T) {
	parseArgs(t)
	client := NewClient(
		WithUrl(UrlTest),
		WithAppId(*appId),
		WithAppSecret(*appSecret),
		WithErp(*erp),
		WithErpVersion(*erpVersion),
	)
	req := NewOrderDetailRequest()
	req.AdorderId = ""
	req.WithWaybill = 1
	resp, err := client.OrderDetail(req)
	if err != nil {
		t.Fatal(".OrderDetail error:", err.Error())
	}
	if resp.Code != 0 {
		t.Fatal(".OrderDetail resp fail:", resp.Message)
	}
}
