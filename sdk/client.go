package sdk

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/go-querystring/query"
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
	url := client.options.Url + request.Path()
	log.Println("url:", url)
	method := request.Method()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Print("req:", req)

	// 如何拼接动态结构体的键值
	v, _ := query.Values(request)
	log.Print("v:", v)
	req.URL.RawQuery = v.Encode()
	log.Println("url string:", req.URL.String())
	log.Println("rawQueury:", req.URL.RawQuery)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	log.Print("do 1:")
	res, err := client.httpClient.Do(req)
	log.Print("do 2:", err)
	if err != nil {
		log.Println("do error:", err)
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("ioutil.ReadAll(res.Body) error:", err)
		return err
	}
	log.Println("body:", string(body))
	if err = json.Unmarshal(
		body,     // data []byte,
		response, // v interface{}
	); err != nil {
		log.Println("json.Unmarshal() error:", err)
		return err
	}
	return
}
