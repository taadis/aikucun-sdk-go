package sdk

import (
	"testing"
)

func TestSkuList(t *testing.T) {
	parseArgs(t)
	client := NewClient(
		WithUrl(UrlTest),
		WithAppId(*appId),
		WithAppSecret(*appSecret),
		WithErp(*erp),
		WithErpVersion(*erpVersion),
	)
	req := &SkuListRequest{
		ActivityId: "",
		Page:       1,
		PageSize:   2,
	}
	resp, err := client.SkuList(req)
	if err != nil {
		t.Fatal(".SkuList() error:", err.Error())
	}
	t.Log(resp)
}
