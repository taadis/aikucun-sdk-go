package sdk

import (
	"encoding/json"
	"io/ioutil"
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
	method := request.Method()
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	// 如何拼接动态结构体的键值
	v, _ := query.Values(request)
	req.URL.RawQuery = v.Encode()
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(
		body,     // data []byte,
		response, // v interface{}
	); err != nil {
		return err
	}
	return
}
