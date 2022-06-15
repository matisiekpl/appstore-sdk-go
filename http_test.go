package appstore

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type HttpRequestBuilderTestSuite struct {
	suite.Suite
	cfg      *Config
	token    *AuthToken
	ctx      context.Context
	testable *RequestBuilder
}

func (suite *HttpRequestBuilderTestSuite) SetupTest() {
	suite.cfg = buildStubConfig()
	suite.token = buildStubAuthToken()
	suite.ctx = context.Background()
	suite.testable = &RequestBuilder{cfg: suite.cfg, token: suite.token}
}

func (suite *HttpRequestBuilderTestSuite) TestBuildHeaders() {
	headers := suite.testable.buildHeaders()
	assert.Equal(suite.T(), "application/a-gzip", headers.Get("Accept"))
	assert.Equal(suite.T(), "gzip", headers.Get("Accept-Encoding"))
	assert.Equal(suite.T(), "Bearer "+suite.token.Token, headers.Get("Authorization"))
}

func (suite *HttpRequestBuilderTestSuite) TestIsValidTokenSuccess() {
	suite.token.ExpiresAt = time.Now().Unix() + 1000
	assert.True(suite.T(), suite.testable.isValidToken())
}

func (suite *HttpRequestBuilderTestSuite) TestIsValidTokenExpired() {
	suite.token.ExpiresAt = time.Now().Unix() - 1000
	assert.False(suite.T(), suite.testable.isValidToken())
}

func (suite *HttpRequestBuilderTestSuite) TestBuildUriWithoutQueryParams() {
	uri, err := suite.testable.buildUri("qwerty", nil)
	assert.NotEmpty(suite.T(), uri)
	assert.Equal(suite.T(), "https://github.com/qwerty", uri.String())
	assert.Nil(suite.T(), err)
}

func (suite *HttpRequestBuilderTestSuite) TestBuildUriWithQueryParams() {
	data := make(map[string]interface{})
	data["foo"] = "bar"
	data["bar"] = "baz"

	uri, err := suite.testable.buildUri("qwerty", data)
	assert.NotEmpty(suite.T(), uri)
	assert.Equal(suite.T(), "https://github.com/qwerty?bar=baz&foo=bar", uri.String())
	assert.Nil(suite.T(), err)
}

func (suite *HttpRequestBuilderTestSuite) TestBuildRequestGET() {
	result, err := suite.testable.BuildRequest(suite.ctx, "get", "foo", map[string]interface{}{"foo": "bar"}, map[string]interface{}{"foo": "bar"})
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.MethodGet, result.Method)
	assert.Equal(suite.T(), "https://github.com/foo?foo=bar", result.URL.String())
	assert.Equal(suite.T(), "application/a-gzip", result.Header.Get("Accept"))
	assert.Equal(suite.T(), "gzip", result.Header.Get("Accept-Encoding"))
	assert.Equal(suite.T(), "Bearer "+suite.token.Token, result.Header.Get("Authorization"))
	assert.Nil(suite.T(), result.Body)
}

func TestHttpRequestBuilderTestSuite(t *testing.T) {
	suite.Run(t, new(HttpRequestBuilderTestSuite))
}

type HttpNewResponseHandlerTestSuite struct {
	suite.Suite
}

func (suite *HttpNewResponseHandlerTestSuite) TestResponseContentTypeGzip() {
	result := NewResponseHandler(ResponseContentTypeGzip, false)
	assert.Implements(suite.T(), (*ResponseHandlerInterface)(nil), result)
	assert.IsType(suite.T(), (*ResponseHandlerGzip)(nil), result)
}

func (suite *HttpNewResponseHandlerTestSuite) TestResponseHandlerJson() {
	result := NewResponseHandler(ResponseContentTypeJson, false)
	assert.Implements(suite.T(), (*ResponseHandlerInterface)(nil), result)
	assert.IsType(suite.T(), (*ResponseHandlerJson)(nil), result)
}

func (suite *HttpNewResponseHandlerTestSuite) TestByDefault() {
	result := NewResponseHandler("foo", false)
	assert.Implements(suite.T(), (*ResponseHandlerInterface)(nil), result)
	assert.IsType(suite.T(), (*ResponseHandlerJson)(nil), result)
}

func TestHttpNewResponseHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(HttpNewResponseHandlerTestSuite))
}

type HttpResponseHandlerJsonTestSuite struct {
	suite.Suite
	testable *ResponseHandlerJson
}

func (suite *HttpResponseHandlerJsonTestSuite) SetupTest() {
	suite.testable = &ResponseHandlerJson{}
}

func (suite *HttpResponseHandlerJsonTestSuite) TestUnmarshalBody() {
	data, _ := ioutil.ReadFile("stubs/errors/invalid.parameter.json")
	var resp ResponseBody
	err := suite.testable.UnmarshalBody(data, &resp)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "foo", resp.Errors[0].Id)
	assert.Equal(suite.T(), "400", resp.Errors[0].Status)
	assert.Equal(suite.T(), "PARAMETER_ERROR.INVALID", resp.Errors[0].Code)
	assert.Equal(suite.T(), "A parameter has an invalid value", resp.Errors[0].Title)
	assert.Equal(suite.T(), "The version parameter you have specified is invalid. The latest version for this report is 1_0.", resp.Errors[0].Detail)
}

func (suite *HttpResponseHandlerJsonTestSuite) TestReadBody() {
	expected, _ := ioutil.ReadFile("stubs/errors/invalid.parameter.json")
	resp := buildStubResponseFromFile(http.StatusOK, "stubs/errors/invalid.parameter.json")
	data, err := suite.testable.ReadBody(resp)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), data)
	assert.Equal(suite.T(), expected, data)
	assert.Empty(suite.T(), body)
}

func (suite *HttpResponseHandlerJsonTestSuite) TestRestoreBody() {
	expected, _ := ioutil.ReadFile("stubs/errors/invalid.parameter.json")
	resp := buildStubResponseFromFile(http.StatusOK, "stubs/errors/invalid.parameter.json")
	data, _ := suite.testable.ReadBody(resp)
	closer, err := suite.testable.RestoreBody(data)
	resp.Body = closer

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expected, data)
	assert.NotEmpty(suite.T(), body)
	assert.Equal(suite.T(), expected, body)
}

func TestHttpResponseHandlerJsonTestSuite(t *testing.T) {
	suite.Run(t, new(HttpResponseHandlerJsonTestSuite))
}

type HttpResponseHandlerGzipTestSuite struct {
	suite.Suite
	testable *ResponseHandlerGzip
}

func (suite *HttpResponseHandlerGzipTestSuite) SetupTest() {
	suite.testable = &ResponseHandlerGzip{}
}

func (suite *HttpResponseHandlerGzipTestSuite) TestUnmarshalBody() {
	reportData, _ := ioutil.ReadFile("stubs/reports/sales/sales.tsv")
	reports := []*SalesReport{}
	err := suite.testable.UnmarshalBody(reportData, &reports)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1234567890, reports[0].AppleIdentifier.Value())
	assert.Equal(suite.T(), "2020-10-05", reports[0].BeginDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), "2020-10-05", reports[0].EndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), float64(299), reports[0].CustomerPrice.Value())
	assert.Equal(suite.T(), 209.3000030517578, reports[0].DeveloperProceeds.Value())
	assert.Equal(suite.T(), float64(12), reports[0].Units.Value())
}

func (suite *HttpResponseHandlerGzipTestSuite) TestReadBody() {
	expected, _ := ioutil.ReadFile("stubs/reports/sales/sales.tsv")
	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/sales.tsv")
	data, err := suite.testable.ReadBody(resp)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), data)
	assert.Equal(suite.T(), expected, data)
	assert.Empty(suite.T(), body)
}

func (suite *HttpResponseHandlerGzipTestSuite) TestRestoreBody() {
	expectedRaw, _ := ioutil.ReadFile("stubs/reports/sales/sales.tsv")
	//expectedGzipped, _ := loadStubResponseDataGzipped("stubs/reports/sales/sales.tsv")
	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/sales.tsv")
	data, err := suite.testable.ReadBody(resp)
	closer, err := suite.testable.RestoreBody(data)
	resp.Body = closer

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRaw, data)
	assert.NotEmpty(suite.T(), body)
	//fmt.Println(expectedGzipped)
	//fmt.Println(body)
	//assert.Equal(t, expectedGzipped, body)
}

func TestHttpResponseHandlerGzipTestSuite(t *testing.T) {
	suite.Run(t, new(HttpResponseHandlerGzipTestSuite))
}

type HttpTransportTestSuite struct {
	suite.Suite
	cfg      *Config
	token    *AuthToken
	ctx      context.Context
	testable *Transport
}

func (suite *HttpTransportTestSuite) SetupTest() {
	suite.cfg = buildStubConfig()
	suite.token = buildStubAuthToken()
	suite.ctx = context.Background()
	suite.testable = NewHttpTransport(suite.cfg, suite.token, nil)
	httpmock.Activate()
}

func (suite *HttpTransportTestSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (suite *HttpTransportTestSuite) TestNewHttpTransport() {
	assert.NotEmpty(suite.T(), suite.testable)
	assert.NotEmpty(suite.T(), suite.testable.http)
	assert.NotEmpty(suite.T(), suite.testable.rb)
}

func (suite *HttpTransportTestSuite) TestRequestSuccess() {
	body, _ := loadStubResponseData("stubs/reports/sales/sales.tsv")

	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, err := suite.testable.SendRequest(suite.ctx, "GET", "foo", nil, nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
}

func (suite *HttpTransportTestSuite) TestRequestGETSuccess() {
	body, _ := loadStubResponseData("stubs/reports/sales/sales.tsv")

	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, err := suite.testable.Get(suite.ctx, "foo", nil)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
}

func (suite *HttpTransportTestSuite) TestRequestGETInvalidToken() {
	body, _ := loadStubResponseData("stubs/reports/sales/sales.tsv")

	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	suite.testable.rb.token.ExpiresAt = time.Now().Unix() - 1000
	resp, err := suite.testable.Get(suite.ctx, "foo", nil)
	assert.Nil(suite.T(), resp)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "transport.request invalid token: <nil>", err.Error())
}

func TestHttpTransportTestSuite(t *testing.T) {
	suite.Run(t, new(HttpTransportTestSuite))
}

type HttpTestSuite struct {
	suite.Suite
}

func (suite *HttpTestSuite) TestNewDefaultHttpClient() {
	client := NewDefaultHttpClient()
	assert.NotEmpty(suite.T(), client)
	assert.NotEmpty(suite.T(), client.Transport)
}

func (suite *HttpTestSuite) TestResponseBodyIsSuccess() {
	rsp := &ResponseBody{status: http.StatusAccepted}
	assert.True(suite.T(), rsp.IsSuccess())
	rsp.status = http.StatusMultipleChoices
	assert.False(suite.T(), rsp.IsSuccess())
	rsp.status = http.StatusBadRequest
	assert.False(suite.T(), rsp.IsSuccess())
}

func (suite *HttpTestSuite) TestResponseBodyGetError() {
	rsp := &ResponseBody{}
	assert.Equal(suite.T(), "", rsp.GetError())
	handler := &ResponseHandlerJson{}
	data, _ := ioutil.ReadFile("stubs/errors/invalid.parameter.json")
	_ = handler.UnmarshalBody(data, &rsp)
	assert.Equal(suite.T(), "The version parameter you have specified is invalid. The latest version for this report is 1_0.", rsp.GetError())
}

func TestHttpTestSuite(t *testing.T) {
	suite.Run(t, new(HttpTestSuite))
}
