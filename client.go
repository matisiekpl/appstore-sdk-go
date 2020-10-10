package appstore_sdk

import "net/http"

type Client struct {
	transport *Transport
	Cfg       *Config
}

func NewClientFromConfig(cfg *Config, token *AuthToken, cl *http.Client) *Client {
	if cl == nil {
		cl = NewDefaultHttpClient()
	}
	client := &Client{Cfg: cfg, transport: NewHttpTransport(cfg, token, cl)}
	return client
}
