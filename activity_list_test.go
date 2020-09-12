package aikucun

import (
	"testing"
)

func TestActivityListResponseString(t *testing.T) {
	resp := &ActivityListResponse{}
	s := resp.String()
	t.Log(s)
	if s == "" {
		t.Fatal(s)
	}
}

func TestActivityList(t *testing.T) {
	parseArgs(t)

	url := ""
	if *isPro {
		url = UrlPro
	} else {
		url = UrlTest
	}
	client := NewClient(
		WithUrl(url),
		WithAppId(*appId),
		WithAppSecret(*appSecret),
		WithErp(*erp),
		WithErpVersion(*erpVersion),
	)
	req := NewActivityListRequest()
	req.Status = 1
	t.Log("req:", req)
	resp, err := client.ActivityList(req)
	if err != nil {
		t.Fatal(err.Error())
	}
	// body: {"code":10003,"data":{},"message":"参数错误","status":"error"}
	if resp.Code != 0 {
		t.Fatal(resp)
	}
}
