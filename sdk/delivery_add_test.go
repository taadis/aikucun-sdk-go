package sdk

import (
	"testing"

	"gitee.com/taadis/aikucun-sdk-go/conf"
)

// TestDeliveryAdd
func TestDeliveryAdd(t *testing.T) {
	parseArgs(t)
	client := NewClient(
		WithUrl(conf.UrlTest),
		WithAppId(*appId),
		WithAppSecret(*appSecret),
		WithErp(*erp),
		WithErpVersion(*erpVersion),
	)
	req := &DeliveryAddRequest{
		AdorderId: "xxx",
	}
	resp, err := client.DeliveryAdd(req)
	if err != nil {
		t.Fatal(".DeliveryAdd error:", err.Error())
	}
	if resp.Code != 0 {
		t.Fatal(".DeliveryAdd resp fail:", resp.Message)
	}
}
