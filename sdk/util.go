package sdk

import (
	"crypto/sha1"
	"fmt"
	"log"
	"sort"
)

// GetSign
func GetSign(appId string, appSecret string, nonceStr string, erp string, erpVersion string, timestamp string, queryParams map[string]interface{}, postJson string) (string, error) {
	obj := make(map[string]interface{})
	obj["appid"] = appId
	obj["appsecret"] = appSecret
	obj["noncestr"] = nonceStr
	obj["timestamp"] = timestamp //time.Now().Unix()
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
			str += fmt.Sprint("&", key, "=", obj[key].(string))
		} else {
			str += fmt.Sprint(key, "=", obj[key].(string))
		}
	}
	log.Println("加密前的字符串:", str)

	// sha1
	ret := fmt.Sprintf("%x", sha1.New().Sum([]byte(str)))
	return ret, nil
}
