package sdk

import (
	"testing"

	"gitee.com/taadis/aikucun-sdk-go/conf"
)

func TestActivityList(t *testing.T) {
	client := NewClient(WithUrl(conf.TestUrl), WithAppId(""), WithAppSecret(""))
	res, err := client.ActivityList(nil)
	t.Log("err:", err.Error())
	if err != nil {
		t.Log(err.Error())
		t.FailNow()
	}
	t.Log(res)
}
