package aikucun

import (
	"testing"
)

//
func TestOrderListAllResponseString(t *testing.T) {
	resp := &OrderListAllResponse{}
	s := resp.String()
	t.Log(s)
}

// TestOrderListAll
func TestOrderListAll(t *testing.T) {
	parseArgs(t)
	client := NewClient(
		WithUrl(UrlTest),
		WithAppId(*appId),
		WithAppSecret(*appSecret),
		WithErp(*erp),
		WithErpVersion(*erpVersion),
	)
	req := NewOrderListAllRequest()
	req.ActivityId = ""
	req.Page = 1
	req.PageSize = 2
	req.WithWaybill = 1

	resp, err := client.OrderListAll(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	if resp.Code != 0 {
		t.Fatal(resp.Message)
	}
	t.Log("resp:", resp)
}
