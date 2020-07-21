package sdk

import (
	"crypto/sha1"
	"encoding/hex"
	"log"
	"sort"
	"strconv"
	"strings"
)

// GetSign
func GetSign(appId string, appSecret string, nonceStr string, erp string, erpVersion string, timestamp string, queryParams map[string]interface{}, postJson string) (string, error) {
	obj := make(map[string]interface{})
	obj["appid"] = appId
	obj["appsecret"] = appSecret
	obj["noncestr"] = nonceStr
	obj["timestamp"] = timestamp
	obj["erp"] = erp
	obj["erpversion"] = erpVersion
	log.Println("obj:", obj)

	// add queryParams
	if queryParams != nil {
		log.Println("queryParams != nil 添加 GET 查询参数")
		for key, value := range queryParams {
			obj[key] = value
		}
		log.Println("添加 GET 查询参数后的对象为:", obj)
	}

	// add postJson
	if postJson != "" {
		log.Println("postJson 不为空, 准备添加 POST 查询参数")
		obj["body"] = postJson
		log.Println("添加 POST 查询参数后的对象为:", obj)
	}

	// sort
	var keys []string
	for key := range obj {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var sb strings.Builder
	for index, key := range keys {
		if index != 0 {
			sb.WriteString("&")
		}
		sb.WriteString(key)
		sb.WriteString("=")
		switch obj[key].(type) {
		case string:
			sb.WriteString(obj[key].(string))
		case int:
			sb.WriteString(strconv.Itoa(obj[key].(int)))
		}
	}
	log.Println("加密前的字符串:", sb.String())

	// sha1
	ret := Sha1Sum(sb.String())
	log.Println("加密后的sign:", ret)
	return ret, nil
}

// Sha1Sum
func Sha1Sum(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
