package sdk

import (
	"encoding/json"
	"log"
	"time"
)

// GetSign
func GetSign(appId string, appSecret string, nonceStr string, erp string, erpVersion string, timestamp string) (string, error) {
	obj := make(map[string]interface{})
	obj["appid"] = appId
	obj["appsecret"] = appSecret
	obj["noncestr"] = nonceStr
	obj["timestamp"] = time.Now().Unix()
	obj["erp"] = erp
	obj["erpversion"] = erpVersion
	log.Println(obj)

	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
