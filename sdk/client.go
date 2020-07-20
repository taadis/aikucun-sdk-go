package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
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
	postJson, err := json.Marshal(request)
	log.Println("postJoson:", string(postJson))
	if request.Method() == http.MethodGet {
		for k, v := range queryParams {
			m[k] = v
		}
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
		resp, err = http.DefaultClient.Do(req)
	} else {
		return errors.New("method not allowed")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Println("body:", string(body))

	if err = json.Unmarshal(
		body,     // data []byte,
		response, // v interface{}
	); err != nil {
		log.Println("json.Unmarshal() error:", err.Error())
		return err
	}
	return
}
