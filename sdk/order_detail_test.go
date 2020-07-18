package sdk

import (
	"testing"

	"gitee.com/taadis/aikucun-sdk-go/conf"
)

// TestOrderDetail
func TestOrderDetail(t *testing.T) {
	parseArgs(t)
	client := NewClient(
		WithUrl(conf.UrlTest),
		WithAppId(*appId),
		WithAppSecret(*appSecret),
		WithErp(*erp),
		WithErpVersion(*erpVersion),
	)
	req := &OrderDetailRequest{
		AdorderId:   "",
		WithWaybill: 1,
	}
	resp, err := client.OrderDetail(req)
	if err != nil {
		t.Fatal(".OrderDetail error:", err.Error())
	}
	if resp.Code != 0 {
		t.Fatal(".OrderDetail resp fail:", resp.Message)
	}
}
