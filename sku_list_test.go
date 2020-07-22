package aikucun

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
	req := NewSkuListRequest()
	req.ActivityId = ""
	req.Page = 1
	req.PageSize = 2
	resp, err := client.SkuList(req)
	if err != nil {
		t.Fatal(".SkuList() error:", err.Error())
	}
	t.Log(resp)
}
