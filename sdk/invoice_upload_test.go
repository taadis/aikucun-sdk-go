package sdk

import (
	"testing"

	"gitee.com/taadis/aikucun-sdk-go/conf"
)

// TestInvoiceUpload
func TestInvoiceUpload(t *testing.T) {
	parseArgs(t)
	client := NewClient(
		WithUrl(conf.UrlTest),
		WithAppId(*appId),
		WithAppSecret(*appSecret),
		WithErp(*erp),
		WithErpVersion(*erpVersion),
	)
	req := &InvoiceUploadRequest{}
	resp, err := client.InvoiceUpload(req)
	if err != nil {
		t.Fatal(".InvoiceUpload error:", err.Error())
	}
	if resp.Code != 0 {
		t.Fatal(".InvoiceUpload resp fail:", resp.Message)
	}
}
