package appstore

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func Test_Sales_GetReports_InvalidFilter(t *testing.T) {
	config := buildStubConfig()
	transport := buildStubHttpTransport()
	resource := &SalesReportsResource{ResourceAbstract: newResourceAbstract(transport, config)}
	filter := &SalesReportsBaseFilter{}
	filter.TypeSales().SubTypeSummary().Version10()
	ctx := context.Background()
	_, err := resource.GetReports(ctx, filter)
	assert.Error(t, err)
	assert.Equal(t, "SalesReportsResource.GetReports invalid filter: SalesReportsBaseFilter.IsValid: Frequency is required", err.Error())
}

func Test_Sales_GetReports_Success(t *testing.T) {
	config := buildStubConfig()
	transport := buildStubHttpTransport()

	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/sales.tsv")
	httpmock.RegisterResponder("GET", config.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(resp))

	resource := &SalesReportsResource{ResourceAbstract: newResourceAbstract(transport, config)}
	filter := &SalesReportsBaseFilter{}
	filter.TypeSales().SubTypeSummary().Version10().Daily()
	ctx := context.Background()
	_, err := resource.GetReports(ctx, filter)
	assert.NoError(t, err)
}

//func Test_Sales_GetSalesReports_Success(t *testing.T) {
//	config := buildStubConfig()
//	transport := buildStubHttpTransport()
//
//	rsp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/sales.tsv")
//	rsp.Header.Set("Content-Type", ResponseContentTypeGzip)
//	httpmock.RegisterResponder("GET", config.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))
//
//	resource := &SalesReportsResource{ResourceAbstract: newResourceAbstract(transport, config)}
//	filter := NewSalesReportsFilter()
//	filter.SubTypeSummary().Version10().Daily()
//	ctx := context.Background()
//	result, resp, err := resource.GetSalesReports(ctx, filter)
//	assert.NoError(t, err)
//	assert.NotEmpty(t, resp)
//	assert.NotEmpty(t, result)
//}

func Test_Sales_GetReport_BuildQueryParams(t *testing.T) {
	config := buildStubConfig()
	token := buildStubAuthToken()
	transport := NewHttpTransport(config, token, nil)
	resource := &SalesReportsResource{ResourceAbstract: newResourceAbstract(transport, config)}
	filter := &SalesReportsBaseFilter{}
	filter.TypeSales().SubTypeSummary().Version10().Daily()
	result := resource.buildQueryParams(filter)
	qs := make(map[string]interface{})
	qs["filter[reportSubType]"] = string(SalesReportSubTypeSummary)
	qs["filter[reportType]"] = string(SalesReportTypeSales)
	qs["filter[frequency]"] = string(SalesReportFrequencyDaily)
	qs["filter[version]"] = string(SalesReportVersion10)
	qs["filter[vendorNumber]"] = config.VendorNo
	assert.Equal(t, qs, result)
}
