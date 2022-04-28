package appstore

import (
	"context"
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Test_HTTP_RequestBuilder_BuildHeaders(t *testing.T) {
	cfg := buildStubConfig()
	token := buildStubAuthToken()
	builder := RequestBuilder{cfg: cfg, token: token}
	headers := builder.buildHeaders()
	assert.Equal(t, "application/a-gzip", headers.Get("Accept"))
	assert.Equal(t, "gzip", headers.Get("Accept-Encoding"))
	assert.Equal(t, "Bearer "+token.Token, headers.Get("Authorization"))
}

func Test_HTTP_RequestBuilder_IsValidTokenSuccess(t *testing.T) {
	cfg := buildStubConfig()
	token := buildStubAuthToken()
	token.ExpiresAt = time.Now().Unix() + 1000
	builder := RequestBuilder{cfg: cfg, token: token}
	assert.True(t, builder.isValidToken())
}

func Test_HTTP_RequestBuilder_IsValidTokenExpired(t *testing.T) {
	cfg := buildStubConfig()
	token := buildStubAuthToken()
	token.ExpiresAt = time.Now().Unix() - 1000
	builder := RequestBuilder{cfg: cfg, token: token}
	assert.False(t, builder.isValidToken())
}

func Test_HTTP_RequestBuilder_BuildUriWithoutQueryParams(t *testing.T) {
	cfg := buildStubConfig()
	builder := RequestBuilder{cfg: cfg}
	uri, err := builder.buildUri("qwerty", nil)
	assert.NotEmpty(t, uri)
	assert.Equal(t, "https://github.com/qwerty", uri.String())
	assert.Nil(t, err)
}

func Test_HTTP_RequestBuilder_BuildUriWithQueryParams(t *testing.T) {
	cfg := buildStubConfig()
	builder := RequestBuilder{cfg: cfg}

	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	uri, err := builder.buildUri("qwerty", data)
	assert.NotEmpty(t, uri)
	assert.Equal(t, "https://github.com/qwerty?bar=baz&foo=bar", uri.String())
	assert.Nil(t, err)
}

func Test_HTTP_RequestBuilder_BuildRequestGET(t *testing.T) {
	cfg := buildStubConfig()
	token := buildStubAuthToken()
	builder := RequestBuilder{cfg: cfg, token: token}

	ctx := context.Background()
	result, err := builder.BuildRequest(ctx, "get", "foo", map[string]interface{}{"foo": "bar"}, map[string]interface{}{"foo": "bar"})
	assert.NoError(t, err)
	assert.Equal(t, http.MethodGet, result.Method)
	assert.Equal(t, "https://github.com/foo?foo=bar", result.URL.String())
	assert.Equal(t, "application/a-gzip", result.Header.Get("Accept"))
	assert.Equal(t, "gzip", result.Header.Get("Accept-Encoding"))
	assert.Equal(t, "Bearer "+token.Token, result.Header.Get("Authorization"))
	assert.Nil(t, result.Body)
}

func Test_HTTP_NewResponseHandler_ResponseContentTypeGzip(t *testing.T) {
	result := NewResponseHandler(ResponseContentTypeGzip)
	assert.Implements(t, (*ResponseHandlerInterface)(nil), result)
	assert.IsType(t, (*ResponseHandlerGzip)(nil), result)
}

func Test_HTTP_NewResponseHandler_ResponseHandlerJson(t *testing.T) {
	result := NewResponseHandler(ResponseContentTypeJson)
	assert.Implements(t, (*ResponseHandlerInterface)(nil), result)
	assert.IsType(t, (*ResponseHandlerJson)(nil), result)
}

func Test_HTTP_NewResponseHandler_ByDefault(t *testing.T) {
	result := NewResponseHandler("foo")
	assert.Implements(t, (*ResponseHandlerInterface)(nil), result)
	assert.IsType(t, (*ResponseHandlerJson)(nil), result)
}

func Test_HTTP_ResponseHandlerJson_UnmarshalBody(t *testing.T) {
	handler := &ResponseHandlerJson{}
	data, _ := ioutil.ReadFile("stubs/errors/invalid.parameter.json")
	var resp ResponseBody
	err := handler.UnmarshalBody(data, &resp)
	assert.NoError(t, err)
	assert.Equal(t, "foo", resp.Errors[0].Id)
	assert.Equal(t, "400", resp.Errors[0].Status)
	assert.Equal(t, "PARAMETER_ERROR.INVALID", resp.Errors[0].Code)
	assert.Equal(t, "A parameter has an invalid value", resp.Errors[0].Title)
	assert.Equal(t, "The version parameter you have specified is invalid. The latest version for this report is 1_0.", resp.Errors[0].Detail)
}

func Test_HTTP_ResponseHandlerJson_ReadBody(t *testing.T) {
	handler := &ResponseHandlerJson{}
	expected, _ := ioutil.ReadFile("stubs/errors/invalid.parameter.json")
	resp := buildStubResponseFromFile(http.StatusOK, "stubs/errors/invalid.parameter.json")
	data, err := handler.ReadBody(resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
	assert.Equal(t, expected, data)
	assert.Empty(t, body)
}

func Test_HTTP_ResponseHandlerJson_RestoreBody(t *testing.T) {
	handler := &ResponseHandlerJson{}
	expected, _ := ioutil.ReadFile("stubs/errors/invalid.parameter.json")
	resp := buildStubResponseFromFile(http.StatusOK, "stubs/errors/invalid.parameter.json")
	data, _ := handler.ReadBody(resp)
	closer, err := handler.RestoreBody(data)
	resp.Body = closer

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, expected, data)
	assert.NotEmpty(t, body)
	assert.Equal(t, expected, body)
}

func Test_HTTP_ResponseHandlerGzip_UnmarshalBody(t *testing.T) {
	handler := &ResponseHandlerGzip{}
	reportData, _ := ioutil.ReadFile("stubs/reports/sales/sales.tsv")
	reports := []*SalesReport{}
	err := handler.UnmarshalBody(reportData, &reports)
	assert.NoError(t, err)
	assert.Equal(t, 1234567890, reports[0].AppleIdentifier.Value())
	assert.Equal(t, "2020-10-05", reports[0].BeginDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, "2020-10-05", reports[0].EndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, float64(299), reports[0].CustomerPrice.Value())
	assert.Equal(t, 209.3000030517578, reports[0].DeveloperProceeds.Value())
	assert.Equal(t, float64(12), reports[0].Units.Value())
}

func Test_HTTP_ResponseHandlerGzip_ReadBody(t *testing.T) {
	handler := &ResponseHandlerGzip{}
	expected, _ := ioutil.ReadFile("stubs/reports/sales/sales.tsv")
	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/sales.tsv")
	data, err := handler.ReadBody(resp)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
	assert.Equal(t, expected, data)
	assert.Empty(t, body)
}

func Test_HTTP_ResponseHandlerGzip_RestoreBody(t *testing.T) {
	handler := &ResponseHandlerGzip{}
	expectedRaw, _ := ioutil.ReadFile("stubs/reports/sales/sales.tsv")
	expectedGzipped, _ := loadStubResponseDataGzipped("stubs/reports/sales/sales.tsv")
	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/sales.tsv")
	data, err := handler.ReadBody(resp)
	closer, err := handler.RestoreBody(data)
	resp.Body = closer

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, expectedRaw, data)
	assert.NotEmpty(t, body)
	fmt.Println(expectedGzipped)
	//fmt.Println(body)
	//assert.Equal(t, expectedGzipped, body)
}

func Test_HTTP_Transport_RequestSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := buildStubConfig()
	transport := buildStubHttpTransport()

	body, _ := loadStubResponseData("stubs/reports/sales/sales.tsv")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	resp, _ := transport.SendRequest(ctx, "GET", "foo", nil, nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Transport_RequestGETSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := buildStubConfig()
	transport := buildStubHttpTransport()

	body, _ := loadStubResponseData("stubs/reports/sales/sales.tsv")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	resp, _ := transport.Get(ctx, "foo", nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Transport_RequestGETInvalidToken(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := buildStubConfig()
	transport := buildStubHttpTransport()
	transport.rb.token.ExpiresAt = time.Now().Unix() - 10000

	body, _ := loadStubResponseData("stubs/reports/sales/sales.tsv")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	ctx := context.Background()
	_, err := transport.Get(ctx, "foo", nil)
	assert.Error(t, err)
	assert.Equal(t, "transport.request invalid token: <nil>", err.Error())
}

func Test_HTTP_NewHttpTransport(t *testing.T) {
	cfg := buildStubConfig()
	token := buildStubAuthToken()
	transport := NewHttpTransport(cfg, token, nil)
	assert.NotEmpty(t, transport)
	assert.NotEmpty(t, transport.http)
	assert.NotEmpty(t, transport.rb)
}

func Test_HTTP_NewDefaultHttpClient(t *testing.T) {
	client := NewDefaultHttpClient()
	assert.NotEmpty(t, client)
}

func Test_HTTP_ResponseBody_IsSuccess(t *testing.T) {
	rsp := &ResponseBody{status: http.StatusAccepted}
	assert.True(t, rsp.IsSuccess())
	rsp.status = http.StatusMultipleChoices
	assert.False(t, rsp.IsSuccess())
	rsp.status = http.StatusBadRequest
	assert.False(t, rsp.IsSuccess())
}

func Test_HTTP_ResponseBody_GetError(t *testing.T) {
	rsp := &ResponseBody{}
	assert.Equal(t, "", rsp.GetError())
	handler := &ResponseHandlerJson{}
	data, _ := ioutil.ReadFile("stubs/errors/invalid.parameter.json")
	_ = handler.UnmarshalBody(data, &rsp)
	assert.Equal(t, "The version parameter you have specified is invalid. The latest version for this report is 1_0.", rsp.GetError())
}
