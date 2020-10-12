package appstore_sdk

import (
	"fmt"
	"net/http"
)

//Client common
type Client struct {
	transport *Transport
	auth      *TokenBuilder
	http      *http.Client
	Cfg       *Config
}

//Initialization of client
func (cl *Client) Init() error {
	token, err := cl.auth.BuildAuthToken()
	if err != nil {
		return fmt.Errorf("client@init error: %v", err)
	}
	cl.transport = NewHttpTransport(cl.Cfg, token, cl.http)
	return nil
}

//Get SalesReports resource
func (cl *Client) SalesReports() *SalesReportsResource {
	return &SalesReportsResource{ResourceAbstract: newResourceAbstract(cl.transport, cl.Cfg)}
}

//Get FinancialReports resource
func (cl *Client) FinancialReports() *FinancialReportsResource {
	return &FinancialReportsResource{ResourceAbstract: newResourceAbstract(cl.transport, cl.Cfg)}
}

//Create new client from config
func NewClientFromConfig(cfg *Config, cl *http.Client) *Client {
	if cl == nil {
		cl = NewDefaultHttpClient()
	}
	client := &Client{Cfg: cfg, auth: NewTokenBuilder(cfg), http: cl}
	return client
}
