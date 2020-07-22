package sdk

import (
	"testing"
)

// TestOrderListAll
func TestOrderListAll(t *testing.T) {
	parseArgs(t)
	client := NewClient(
		WithUrl(UrlTest),
		WithAppId(*appId),
		WithAppSecret(*appSecret),
		WithErp(*erp),
		WithErpVersion(*erpVersion),
	)
	req := &OrderListAllRequest{
		ActivityId:  "",
		Page:        1,
		PageSize:    2,
		WithWaybill: 1,
	}
	resp, err := client.OrderListAll(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	if resp.Code != 0 {
		t.Fatal(resp.Message)
	}
	t.Log("resp:", resp)
}
