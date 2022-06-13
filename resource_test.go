package appstore

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type ResourceTestSuite struct {
	suite.Suite
}

func (suite *ResourceTestSuite) TestNewResourceAbstract() {
	config := buildStubConfig()
	token := buildStubAuthToken()
	transport := NewHttpTransport(config, token, nil)
	result := newResourceAbstract(transport, config)
	assert.NotEmpty(suite.T(), result)
	assert.NotEmpty(suite.T(), result.config)
	assert.NotEmpty(suite.T(), result.transport)
}

func (suite *ResourceTestSuite) TestUnmarshalResponseGzip() {
	result := buildStubResourceAbstract()
	reports := []*SalesReport{}
	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/sales.tsv")
	resp.Header.Set("Content-Type", ResponseContentTypeGzip)
	err := result.unmarshalResponse(resp, &reports, false)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1234567890, reports[0].AppleIdentifier.Value())
	assert.Equal(suite.T(), "2020-10-05", reports[0].BeginDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), "2020-10-05", reports[0].EndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), float64(299), reports[0].CustomerPrice.Value())
	assert.Equal(suite.T(), 209.3000030517578, reports[0].DeveloperProceeds.Value())
	assert.Equal(suite.T(), float64(12), reports[0].Units.Value())
}

func (suite *ResourceTestSuite) TestUnmarshalResponseJson() {
	result := buildStubResourceAbstract()
	var body ResponseBody
	resp := buildStubResponseFromFile(http.StatusOK, "stubs/errors/invalid.parameter.json")
	resp.Header.Set("Content-Type", ResponseContentTypeJson)
	err := result.unmarshalResponse(resp, &body, false)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "foo", body.Errors[0].Id)
	assert.Equal(suite.T(), "400", body.Errors[0].Status)
	assert.Equal(suite.T(), "PARAMETER_ERROR.INVALID", body.Errors[0].Code)
	assert.Equal(suite.T(), "A parameter has an invalid value", body.Errors[0].Title)
	assert.Equal(suite.T(), "The version parameter you have specified is invalid. The latest version for this report is 1_0.", body.Errors[0].Detail)
}

func TestResourceTestSuite(t *testing.T) {
	suite.Run(t, new(ResourceTestSuite))
}
