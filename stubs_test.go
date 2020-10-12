package appstore_sdk

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const StubAuthKeyPath string = "stubs/auth/keys/AuthKeyStub_4W5TU4DR28.p8"

func buildStubConfig() *Config {
	return &Config{
		Uri:        "https://github.com",
		IssuerId:   "foo",
		KeyId:      "bar",
		VendorNo:   "baz",
		PrivateKey: StubAuthKeyPath,
		Token:      NewTokenConfig(),
	}
}

func buildStubAuthToken() *AuthToken {
	return &AuthToken{
		Token:     "AuthToken",
		ExpiresAt: time.Now().Unix() + 1000,
	}
}

func buildStubHttpTransport() *Transport {
	return NewHttpTransport(buildStubConfig(), buildStubAuthToken(), nil)
}

func loadStubResponseData(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func buildStubResponseFromString(statusCode int, json string) *http.Response {
	body := ioutil.NopCloser(strings.NewReader(json))
	return &http.Response{Body: body, StatusCode: statusCode}
}

func buildStubResponseFromFile(statusCode int, path string) *http.Response {
	data, _ := loadStubResponseData(path)
	body := ioutil.NopCloser(bytes.NewReader(data))
	return &http.Response{Body: body, StatusCode: statusCode}
}
