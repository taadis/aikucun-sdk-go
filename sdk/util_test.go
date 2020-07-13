package sdk

import (
	"testing"
)

// TestGetSign
func TestGetSign(t *testing.T) {
	var appId, appSecret, nonceStr, erp, erpVersion, timestamp, postJson string
	queryParams := make(map[string]string)
	sign, err := GetSign(appId, appSecret, nonceStr, erp, erpVersion, timestamp, queryParams, postJson)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("sign=", sign)
}

// TestSha1Sum
func TestSha1Sum(t *testing.T) {
	s := "123456"
	ret := Sha1Sum(s)
	if ret != "7c4a8d09ca3762af61e59520943dc26494f8941b" {
		t.Fatal("sha1 sum fail")
	}
}
