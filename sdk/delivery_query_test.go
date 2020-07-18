package sdk

import (
	"testing"

	"gitee.com/taadis/aikucun-sdk-go/conf"
)

// TestDeliveryQuery
func TestDeliveryQuery(t *testing.T) {
	parseArgs(t)
	client := NewClient(
		WithUrl(conf.UrlTest),
		WithAppId(*appId),
		WithAppSecret(*appSecret),
		WithErp(*erp),
		WithErpVersion(*erpVersion),
	)
	req := &DeliveryQueryRequest{
		AdorderId: "",
	}
	resp, err := client.DeliveryQuery(req)
	if err != nil {
		t.Fatal(".DeliveryQuery error:", err.Error())
	}
	if resp.Code != 0 {
		t.Fatal(".DeliveryQuery fail:", resp.Message)
	}
}
