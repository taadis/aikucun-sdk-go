package sdk

import (
	"net/http"
)

// 客户端结构体
type Client struct {
	httpClient *http.Client
	url        string
}
