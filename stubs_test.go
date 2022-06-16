package appstore

import (
	"bytes"
	"compress/gzip"
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
	return NewHttpTransport(buildStubConfig(), buildStubAuthToken(), &http.Client{})
}

func loadStubResponseData(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

func loadStubResponseDataGzipped(path string) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	_, err = zw.Write(data)
	err = zw.Close()
	return buf.Bytes(), err
}

func buildStubResponseFromString(statusCode int, str string) *http.Response {
	body := ioutil.NopCloser(strings.NewReader(str))
	return &http.Response{Body: body, StatusCode: statusCode, Header: http.Header{}}
}

func buildStubResponseFromFile(statusCode int, path string) *http.Response {
	data, _ := loadStubResponseData(path)
	body := ioutil.NopCloser(bytes.NewReader(data))
	return &http.Response{Body: body, StatusCode: statusCode, Header: http.Header{}}
}

func buildStubResponseFromGzip(statusCode int, path string) *http.Response {
	data, _ := loadStubResponseDataGzipped(path)
	body := ioutil.NopCloser(bytes.NewReader(data))
	return &http.Response{Body: body, StatusCode: statusCode, Header: http.Header{}}
}

func buildStubResourceAbstract() ResourceAbstract {
	config := buildStubConfig()
	transport := buildStubHttpTransport()
	return newResourceAbstract(transport, config)
}

func buildStubSalesReportsResource() *SalesReportsResource {
	config := buildStubConfig()
	transport := buildStubHttpTransport()
	return &SalesReportsResource{newResourceAbstract(transport, config)}
}

func buildStubFinancesReportsResource() *FinancesReportsResource {
	config := buildStubConfig()
	transport := buildStubHttpTransport()
	return &FinancesReportsResource{newResourceAbstract(transport, config)}
}
