package appstore_sdk

import (
	"net/http"
	"time"
)

func NewHttpClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:    10,
		IdleConnTimeout: 30 * time.Second,
	}
	return &http.Client{Transport: tr}
}
