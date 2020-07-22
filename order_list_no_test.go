package aikucun

import (
	"testing"
)

// TestOrderListNo
func TestOrderListNo(t *testing.T) {
	parseArgs(t)
	client := NewClient(
		WithUrl(UrlTest),
		WithAppId(*appId),
		WithAppSecret(*appSecret),
		WithErp(*erp),
		WithErpVersion(*erpVersion),
	)
	req := NewOrderListNoRequest()
	req.ActivityId = ""
	req.Page = 1
	req.PageSize = 2
	req.FilterDeliver = 0
	resp, err := client.OrderListNo(req)
	if err != nil {
		t.Fatal(".OrderListNo error:", err.Error())
	}
	if resp.Code != 0 {
		t.Fatal(".OrderListNo resp message:", resp.Message)
	}
}
