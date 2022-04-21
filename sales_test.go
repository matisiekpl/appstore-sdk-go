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
	filter := &SalesReportsFilter{}
	filter.TypeSales().SubTypeSummary().Version10()
	ctx := context.Background()
	_, err := resource.GetReport(ctx, filter)
	assert.Error(t, err)
	assert.Equal(t, "SalesReportsResource@GetReport invalid filter: SalesReportsFilter@IsValid: Frequency is required", err.Error())
}
