package sdk

import (
	"testing"
)

// TestGetSign
func TestGetSign(t *testing.T) {
	var appId, appSecret, nonceStr, erp, erpVersion, timestamp, postJson string
	queryParams := make(map[string]interface{})
	sign, err := GetSign(appId, appSecret, nonceStr, erp, erpVersion, timestamp, queryParams, postJson)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log("sign=", sign)
}
