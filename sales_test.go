package appstore_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Sales_GetReport_InvalidFilter(t *testing.T) {
	config := buildStubConfig()
	token := buildStubAuthToken()
	transport := NewHttpTransport(config, token, nil)
	resource := &SalesReportsResource{ResourceAbstract: newResourceAbstract(transport, config)}
	filter := &SalesReportsFilter{}
	filter.Daily().TypeSales().SubTypeSummary().Version10()
	_, err := resource.GetReport(filter)
	assert.Error(t, err)
	assert.Equal(t, "SalesReportsResource@GetReport invalid filter: SalesReportsFilter@IsValid: ReportDate is required", err.Error())
}
