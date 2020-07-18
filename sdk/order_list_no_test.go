package sdk

import (
	"testing"

	"gitee.com/taadis/aikucun-sdk-go/conf"
)

// TestOrderListNo
func TestOrderListNo(t *testing.T) {
	parseArgs(t)
	client := NewClient(
		WithUrl(conf.UrlTest),
		WithAppId(*appId),
		WithAppSecret(*appSecret),
		WithErp(*erp),
		WithErpVersion(*erpVersion),
	)
	req := &OrderListNoRequest{
		ActivityId:    "",
		Page:          1,
		PageSize:      2,
		FilterDeliver: 0,
	}
	resp, err := client.OrderListNo(req)
	if err != nil {
		t.Fatal(".OrderListNo error:", err.Error())
	}
	if resp.Code != 0 {
		t.Fatal(".OrderListNo resp message:", resp.Message)
	}
}
