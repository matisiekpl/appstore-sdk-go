package appstore_sdk

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
)

func BuildStubConfig() *Config {
	return &Config{
		Uri:        "https://github.com",
		IssuerId:   "foo",
		KeyId:      "bar",
		VendorNo:   "baz",
		PrivateKey: "tmp/key.p8",
		Token:      NewTokenConfig(),
	}
}

func BuildStubAuthToken() *AuthToken {
	return &AuthToken{
		Token:     "AuthToken",
		ExpiresAt: 100,
	}
}

func LoadStubResponseData(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func BuildStubResponseFromString(statusCode int, json string) *http.Response {
	body := ioutil.NopCloser(strings.NewReader(json))
	return &http.Response{Body: body, StatusCode: statusCode}
}

func BuildStubResponseFromFile(statusCode int, path string) *http.Response {
	data, _ := LoadStubResponseData(path)
	body := ioutil.NopCloser(bytes.NewReader(data))
	return &http.Response{Body: body, StatusCode: statusCode}
}
