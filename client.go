package appstore_sdk

import "net/http"

type Client struct {
	Cfg *Config
}

func NewClientFromConfig(cfg *Config, cl *http.Client) *Client {
	client := &Client{Cfg: cfg}
	return client
}
