package aikucun

import (
	"testing"
)

func TestWithUrl(t *testing.T) {
	WithUrl("url")
}

func TestWithAppId(t *testing.T) {
	WithAppId("appid")
}

func TestWithAppSecret(t *testing.T) {
	WithAppSecret("appsecret")
}

func TestNewOptions(t *testing.T) {
	options := NewOptions(WithUrl("url"), WithAppId("appid"), WithAppSecret("appsecret"))
	if !(options.Url == "url" && options.AppId == "appid" && options.AppSecret == "appsecret") {
		t.FailNow()
	}
}
