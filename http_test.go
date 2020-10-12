package appstore_sdk

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
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

func Test_HTTP_Transport_RequestSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := buildStubConfig()
	transport := buildStubHttpTransport()

	body, _ := loadStubResponseData("stubs/reports/sales/sales.tsv")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := transport.Request("GET", "foo", nil, nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Transport_RequestGETSuccess(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := buildStubConfig()
	transport := buildStubHttpTransport()

	body, _ := loadStubResponseData("stubs/reports/sales/sales.tsv")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := transport.Get("foo", nil)
	assert.NotEmpty(t, resp)
}

func Test_HTTP_Transport_RequestGETInvalidToken(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := buildStubConfig()
	transport := buildStubHttpTransport()
	transport.rb.token.ExpiresAt = time.Now().Unix() - 1000

	body, _ := loadStubResponseData("stubs/reports/sales/sales.tsv")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	_, err := transport.Get("foo", nil)
	assert.Error(t, err)
	assert.Equal(t, "transport@request invalid token: <nil>", err.Error())
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

func Test_HTTP_Response_IsSuccessTrue(t *testing.T) {
	response := &Response{raw: buildStubResponseFromFile(http.StatusOK, "stubs/reports/sales/sales.tsv")}
	assert.True(t, response.IsSuccess())
}

func Test_HTTP_Response_IsSuccessFalse(t *testing.T) {
	response := &Response{raw: buildStubResponseFromFile(http.StatusBadRequest, "stubs/reports/sales/sales.tsv")}
	assert.False(t, response.IsSuccess())
}

func Test_HTTP_Response_GetRawResponse(t *testing.T) {
	rsp := buildStubResponseFromFile(http.StatusOK, "stubs/reports/sales/sales.tsv")
	response := &Response{raw: rsp}
	raw := response.GetRawResponse()
	assert.NotEmpty(t, raw)
	assert.Equal(t, http.StatusOK, raw.StatusCode)
}

func Test_HTTP_Response_GetRawBodySuccess(t *testing.T) {
	data, _ := loadStubResponseData("stubs/reports/sales/sales.tsv")
	rsp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/sales.tsv")
	response := NewResponse(rsp)
	str, _ := response.GetRawBody()
	assert.Equal(t, string(data), str)
}

func Test_HTTP_Response_GetRawBodyBadRequest(t *testing.T) {
	data, _ := loadStubResponseData("stubs/reports/sales/sales.tsv")
	rsp := buildStubResponseFromFile(http.StatusBadRequest, "stubs/reports/sales/sales.tsv")
	response := NewResponse(rsp)
	str, _ := response.GetRawBody()
	assert.Equal(t, string(data), str)
}

func Test_HTTP_Response_UnmarshalCSV(t *testing.T) {
	rsp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/sales.tsv")
	response := &Response{raw: rsp}
	reports := []*SalesReportSale{}
	_ = response.UnmarshalCSV(&reports)
	assert.Equal(t, 1234567890, reports[0].AppleIdentifier.Value())
	assert.Equal(t, "2020-10-05", reports[0].BeginDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, "2020-10-05", reports[0].EndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, float64(299), reports[0].CustomerPrice.Value())
	assert.Equal(t, 209.3000030517578, reports[0].DeveloperProceeds.Value())
	assert.Equal(t, 12, reports[0].Units.Value())
}

func Test_HTTP_Response_UnmarshalError(t *testing.T) {
	rsp := buildStubResponseFromFile(http.StatusBadRequest, "stubs/errors/invalid.parameter.json")
	response := &Response{raw: rsp}
	var errorResult *ErrorResult
	_ = response.UnmarshalError(&errorResult)
	err := errorResult.GetError()
	assert.Equal(t, "foo", err.Id)
	assert.Equal(t, "400", err.Status)
	assert.Equal(t, "PARAMETER_ERROR.INVALID", err.Code)
	assert.Equal(t, "A parameter has an invalid value", err.Title)
	assert.Equal(t, "The version parameter you have specified is invalid. The latest version for this report is 1_0.", err.Detail)
	assert.Equal(t, "filter[version]", err.Source.Parameter)
}

func Test_HTTP_NewResponse(t *testing.T) {
	rsp := buildStubResponseFromFile(http.StatusOK, "stubs/reports/sales/sales.tsv")
	response := NewResponse(rsp)
	assert.NotEmpty(t, response)
	assert.NotEmpty(t, response.raw)
}
