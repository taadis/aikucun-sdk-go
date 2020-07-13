package sdk

import (
	"testing"

	"gitee.com/taadis/aikucun-sdk-go/conf"
)

func TestActivityList(t *testing.T) {
	c := conf.GetConf()
	client := NewClient(
		WithUrl(c.TestUrl),
		WithAppId(""),
		WithAppSecret(""),
		WithErp(""),
		WithErpVersion(""),
	)
	req := &ActivityListRequest{
		Status: 1,
	}
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
