package sdk

import (
	"testing"

	"gitee.com/taadis/aikucun-sdk-go/conf"
)

func TestSkuList(t *testing.T) {
	parseArgs(t)
	client := NewClient(
		WithUrl(conf.UrlTest),
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
