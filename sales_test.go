package appstore

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sales_GetReport_InvalidFilter(t *testing.T) {
	config := buildStubConfig()
	token := buildStubAuthToken()
	transport := NewHttpTransport(config, token, nil)
	resource := &SalesReportsResource{ResourceAbstract: newResourceAbstract(transport, config)}
	filter := &SalesReportsBaseFilter{}
	filter.TypeSales().SubTypeSummary().Version10()
	ctx := context.Background()
	_, err := resource.GetReports(ctx, filter)
	assert.Error(t, err)
	assert.Equal(t, "SalesReportsResource.GetReports invalid filter: SalesReportsBaseFilter.IsValid: Frequency is required", err.Error())
}

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
