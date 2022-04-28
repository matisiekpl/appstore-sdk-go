package appstore

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_Resource_NewResourceAbstract(t *testing.T) {
	config := buildStubConfig()
	token := buildStubAuthToken()
	transport := NewHttpTransport(config, token, nil)
	result := newResourceAbstract(transport, config)
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.config)
	assert.NotEmpty(t, result.transport)
}

func Test_Resource_UnmarshalResponseGzip(t *testing.T) {
	config := buildStubConfig()
	token := buildStubAuthToken()
	transport := NewHttpTransport(config, token, nil)
	result := newResourceAbstract(transport, config)
	reports := []*SalesReport{}
	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/sales.tsv")
	resp.Header.Set("Content-Type", ResponseContentTypeGzip)
	err := result.unmarshalResponse(resp, &reports)
	assert.NoError(t, err)
	assert.Equal(t, 1234567890, reports[0].AppleIdentifier.Value())
	assert.Equal(t, "2020-10-05", reports[0].BeginDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, "2020-10-05", reports[0].EndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, float64(299), reports[0].CustomerPrice.Value())
	assert.Equal(t, 209.3000030517578, reports[0].DeveloperProceeds.Value())
	assert.Equal(t, float64(12), reports[0].Units.Value())
}

func Test_Resource_UnmarshalResponseJson(t *testing.T) {
	config := buildStubConfig()
	token := buildStubAuthToken()
	transport := NewHttpTransport(config, token, nil)
	result := newResourceAbstract(transport, config)
	var body ResponseBody
	resp := buildStubResponseFromFile(http.StatusOK, "stubs/errors/invalid.parameter.json")
	resp.Header.Set("Content-Type", ResponseContentTypeJson)
	err := result.unmarshalResponse(resp, &body)
	assert.NoError(t, err)
	assert.Equal(t, "foo", body.Errors[0].Id)
	assert.Equal(t, "400", body.Errors[0].Status)
	assert.Equal(t, "PARAMETER_ERROR.INVALID", body.Errors[0].Code)
	assert.Equal(t, "A parameter has an invalid value", body.Errors[0].Title)
	assert.Equal(t, "The version parameter you have specified is invalid. The latest version for this report is 1_0.", body.Errors[0].Detail)
}
