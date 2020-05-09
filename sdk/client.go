package sdk

import (
	"net/http"
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
