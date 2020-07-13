package sdk

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"sort"
)

// GetSign
func GetSign(appId string, appSecret string, nonceStr string, erp string, erpVersion string, timestamp string, queryParams map[string]string, postJson string) (string, error) {
	obj := make(map[string]interface{})
	obj["appid"] = appId
	obj["appsecret"] = appSecret
	obj["noncestr"] = nonceStr
	obj["timestamp"] = timestamp
	obj["erp"] = erp
	obj["erpversion"] = erpVersion
	log.Println(obj)

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
		obj["post"] = postJson
		log.Println("添加 POST 查询参数后的对象为:", obj)
	}

	// sort
	var keys []string
	for key := range obj {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	str := ""
	for _, key := range keys {
		if str != "" {
			str += fmt.Sprint("&", key, "=", obj[key])
		} else {
			str += fmt.Sprint(key, "=", obj[key])
		}
	}
	log.Println("加密前的字符串:", str)

	// sha1
	ret := fmt.Sprintf("%x", sha1.New().Sum([]byte(str)))
	return ret, nil
}

// Sha1Sum
func Sha1Sum(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
