package appstore_sdk

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func Test_HTTP_RequestBuilder_BuildHeaders(t *testing.T) {
	cfg := BuildStubConfig()
	token := BuildStubAuthToken()
	builder := RequestBuilder{cfg: cfg, token: token}
	headers := builder.buildHeaders()
	assert.Equal(t, "application/a-gzip", headers.Get("Accept"))
	assert.Equal(t, "gzip", headers.Get("Accept-Encoding"))
	assert.Equal(t, "Bearer "+token.Token, headers.Get("Authorization"))
}

func Test_HTTP_RequestBuilder_BuildBody(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	body, _ := builder.buildBody(data)
	assert.NotEmpty(t, body)
}

func Test_HTTP_RequestBuilder_IsValidTokenSuccess(t *testing.T) {
	cfg := BuildStubConfig()
	token := BuildStubAuthToken()
	token.ExpiresAt = time.Now().Unix() + 1000
	builder := RequestBuilder{cfg: cfg, token: token}
	assert.True(t, builder.isValidToken())
}

func Test_HTTP_RequestBuilder_IsValidTokenExpired(t *testing.T) {
	cfg := BuildStubConfig()
	token := BuildStubAuthToken()
	token.ExpiresAt = time.Now().Unix() - 1000
	builder := RequestBuilder{cfg: cfg, token: token}
	assert.False(t, builder.isValidToken())
}

func Test_HTTP_RequestBuilder_BuildUriWithoutQueryParams(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}
	uri, err := builder.buildUri("qwerty", nil)
	assert.NotEmpty(t, uri)
	assert.Equal(t, "https://github.com/qwerty", uri.String())
	assert.Nil(t, err)
}

func Test_HTTP_RequestBuilder_BuildUriWithQueryParams(t *testing.T) {
	cfg := BuildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	uri, err := builder.buildUri("qwerty", data)
	assert.NotEmpty(t, uri)
	assert.Equal(t, "https://github.com/qwerty?bar=baz&foo=bar", uri.String())
	assert.Nil(t, err)
}

func Test_HTTP_NewHttpTransport(t *testing.T) {
	cfg := BuildStubConfig()
	token := BuildStubAuthToken()
	transport := NewHttpTransport(cfg, token, nil)
	assert.NotEmpty(t, transport)
	assert.NotEmpty(t, transport.http)
	assert.NotEmpty(t, transport.rb)
}

func Test_HTTP_NewDefaultHttpClient(t *testing.T) {
	client := NewDefaultHttpClient()
	assert.NotEmpty(t, client)
}

func Test_HTTP_Response_IsSuccessTrue(t *testing.T) {
	response := &Response{raw: BuildStubResponseFromFile(http.StatusOK, "stubs/reports/sales/sales.tsv")}
	assert.True(t, response.IsSuccess())
}

func Test_HTTP_Response_IsSuccessFalse(t *testing.T) {
	response := &Response{raw: BuildStubResponseFromFile(http.StatusBadRequest, "stubs/reports/sales/sales.tsv")}
	assert.False(t, response.IsSuccess())
}

func Test_HTTP_Response_GetRawResponse(t *testing.T) {
	rsp := BuildStubResponseFromFile(http.StatusOK, "stubs/reports/sales/sales.tsv")
	response := &Response{raw: rsp}
	raw := response.GetRawResponse()
	assert.NotEmpty(t, raw)
	assert.Equal(t, http.StatusOK, raw.StatusCode)
}

func Test_HTTP_Response_GetRawBody(t *testing.T) {
	data, _ := LoadStubResponseData("stubs/reports/sales/sales.tsv")
	rsp := BuildStubResponseFromFile(http.StatusOK, "stubs/reports/sales/sales.tsv")
	response := &Response{raw: rsp}
	str, _ := response.GetRawBody()
	assert.Equal(t, string(data), str)
}
