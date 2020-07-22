package aikucun

import (
	"testing"
)

// TestInvoiceUpload
func TestInvoiceUpload(t *testing.T) {
	parseArgs(t)
	client := NewClient(
		WithUrl(UrlTest),
		WithAppId(*appId),
		WithAppSecret(*appSecret),
		WithErp(*erp),
		WithErpVersion(*erpVersion),
	)
	req := NewInvoiceUploadRequest()
	resp, err := client.InvoiceUpload(req)
	if err != nil {
		t.Fatal(".InvoiceUpload error:", err.Error())
	}
	if resp.Code != 0 {
		t.Fatal(".InvoiceUpload resp fail:", resp.Message)
	}
}
