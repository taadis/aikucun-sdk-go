package aikucun

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

// 客户端结构体
type Client struct {
	httpClient *http.Client
	options    Options
}

func NewClient(options ...Option) *Client {
	opts := NewOptions(options...)
	client := &Client{
		httpClient: &http.Client{},
		options:    opts,
	}
	return client
}

// Do
func (client *Client) Do(request IRequest, response IResponse) (err error) {

	//
	m := make(map[string]interface{})
	m["appid"] = client.options.AppId
	//m["appsecret"] = client.options.AppSecret
	m["noncestr"] = strconv.FormatInt(time.Now().Unix(), 10)
	m["timestamp"] = strconv.FormatInt(time.Now().Unix(), 10)
	m["erp"] = client.options.Erp
	m["erpversion"] = client.options.ErpVersion
	//m["sign"] = ""

	// get sign
	queryParams := request.Params()
	log.Println("queryParams:", queryParams)
	postJson, err := json.Marshal(request)
	log.Println("postJson:", string(postJson))
	for k, v := range queryParams {
		m[k] = v
	}
	log.Println("m:", m)
	if request.Method() == http.MethodGet {
		sign, err := GetSign(
			m["appid"].(string),
			client.options.AppSecret,
			m["noncestr"].(string),
			m["erp"].(string),
			m["erpversion"].(string),
			m["timestamp"].(string),
			queryParams,
			"")
		if err != nil {
			log.Println("GetSign error:", err.Error())
			return err
		}
		log.Println("sign=", sign)
		m["sign"] = sign
	} else if request.Method() == http.MethodPost {
		m["sign"], err = GetSign(
			m["appid"].(string),
			client.options.AppSecret,
			m["noncestr"].(string),
			m["erp"].(string),
			m["erpversion"].(string),
			m["timestamp"].(string),
			queryParams,
			string(postJson))
		if err != nil {
			log.Println("GetSign error:", err.Error())
			return err
		}
	} else {
		return errors.New("method not allowed")
	}

	//
	u, err := url.Parse(client.options.Url + request.Path())
	if err != nil {
		log.Println("url.Parse error:", err.Error())
		return err
	}
	ps := url.Values{}
	for k, v := range m {
		switch v.(type) {
		case string:
			ps.Set(k, v.(string))
		case int:
			ps.Set(k, strconv.Itoa(v.(int)))
		}
	}
	u.RawQuery = ps.Encode()
	log.Println("url:", u.String())

	//
	var resp *http.Response
	method := request.Method()
	if method == http.MethodGet {
		req, err := http.NewRequest(method, u.String(), nil)
		if err != nil {
			log.Println(err)
			return err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		log.Print("req:", req)
		resp, err = http.DefaultClient.Do(req)
	} else if method == http.MethodPost {
		req, err := http.NewRequest(method, u.String(), bytes.NewBuffer(postJson))
		if err != nil {
			log.Println(err.Error())
			return err
		}
		req.Header.Set("Content-Type", "application/json;charset=utf8")
		resp, err = http.DefaultClient.Do(req)
	} else {
		return errors.New("method not allowed")
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		log.Println("Error json.Decode:", err.Error())
		return err
	}
	return
}
