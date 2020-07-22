package aikucun

import (
	"testing"
)

// TestDeliveryQuery
func TestDeliveryQuery(t *testing.T) {
	parseArgs(t)
	client := NewClient(
		WithUrl(UrlTest),
		WithAppId(*appId),
		WithAppSecret(*appSecret),
		WithErp(*erp),
		WithErpVersion(*erpVersion),
	)
	req := NewDeliveryQueryRequest()
	req.AdorderId = ""
	resp, err := client.DeliveryQuery(req)
	if err != nil {
		t.Fatal(".DeliveryQuery error:", err.Error())
	}
	if resp.Code != 0 {
		t.Fatal(".DeliveryQuery fail:", resp.Message)
	}
}
