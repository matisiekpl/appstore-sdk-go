package appstore

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
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
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := buildStubConfig()
	transport := buildStubHttpTransport()

	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/sales.tsv")
	resp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", config.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(resp))

	resource := &SalesReportsResource{ResourceAbstract: newResourceAbstract(transport, config)}
	filter := &SalesReportsBaseFilter{}
	filter.TypeSales().SubTypeSummary().Version10().Daily()
	ctx := context.Background()
	rsp, err := resource.GetReports(ctx, filter)
	assert.NoError(t, err)
	assert.NotEmpty(t, rsp)
}

func Test_Sales_GetSalesReports_Success(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := buildStubConfig()
	transport := buildStubHttpTransport()
	ar := newResourceAbstract(transport, config)

	rsp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/sales.tsv")
	rsp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", config.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	resource := &SalesReportsResource{ar}
	filter := NewSalesReportsFilter()
	filter.SubTypeSummary().Version10().Daily()
	ctx := context.Background()
	result, resp, err := resource.GetSalesReports(ctx, filter)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)

	assert.True(t, result.IsSuccess())
	assert.Empty(t, result.GetError())
	assert.Empty(t, result.Errors)
	assert.Equal(t, 1234567890, result.Data[0].AppleIdentifier.Value())
	assert.Equal(t, "2020-10-05", result.Data[0].BeginDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, "2020-10-05", result.Data[0].EndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, float64(299), result.Data[0].CustomerPrice.Value())
	assert.Equal(t, 209.3000030517578, result.Data[0].DeveloperProceeds.Value())
	assert.Equal(t, float64(12), result.Data[0].Units.Value())

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(t, body)
}

func Test_Sales_GetSalesReports_Error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := buildStubConfig()
	transport := buildStubHttpTransport()
	ar := newResourceAbstract(transport, config)

	rsp := buildStubResponseFromFile(http.StatusBadRequest, "stubs/errors/invalid.parameter.json")
	rsp.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder("GET", config.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	resource := &SalesReportsResource{ar}
	filter := NewSalesReportsFilter()
	filter.SubTypeSummary().Version10().Daily()
	ctx := context.Background()
	result, resp, err := resource.GetSalesReports(ctx, filter)
	assert.Error(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	assert.Equal(t, "The version parameter you have specified is invalid. The latest version for this report is 1_0.", err.Error())

	assert.False(t, result.IsSuccess())
	assert.Equal(t, "The version parameter you have specified is invalid. The latest version for this report is 1_0.", result.GetError())
	assert.Len(t, result.Errors, 1)
	assert.Empty(t, result.Data)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(t, body)
}

func Test_Sales_GetSubscriptionsReports_Success(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := buildStubConfig()
	transport := buildStubHttpTransport()
	ar := newResourceAbstract(transport, config)

	rsp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/subscriptions.tsv")
	rsp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", config.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	resource := &SalesReportsResource{ar}
	filter := NewSubscriptionsReportsFilter()
	filter.SubTypeSummary().Version12().Daily()
	ctx := context.Background()
	result, resp, err := resource.GetSubscriptionsReports(ctx, filter)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)

	assert.True(t, result.IsSuccess())
	assert.Empty(t, result.GetError())
	assert.Empty(t, result.Errors)
	assert.Equal(t, "FooBarApp", result.Data[0].AppName)
	assert.Equal(t, 1234567890, result.Data[0].AppAppleID.Value())
	assert.Equal(t, "foo.bar.baz", result.Data[0].SubscriptionName)
	assert.Equal(t, 1234567890, result.Data[0].SubscriptionAppleID.Value())
	assert.Equal(t, 1234567890, result.Data[0].SubscriptionGroupID.Value())
	assert.Equal(t, "1 Year", result.Data[0].StandardSubscriptionDuration)
	assert.Equal(t, " ", result.Data[0].PromotionalOfferName)
	assert.Equal(t, " ", result.Data[0].PromotionalOfferID)
	assert.Equal(t, float64(2950), result.Data[0].CustomerPrice.Value())
	assert.Equal(t, "RUB", result.Data[0].CustomerCurrency)
	assert.Equal(t, 2065.00, result.Data[0].DeveloperProceeds.Value())
	assert.Equal(t, "RUB", result.Data[0].ProceedsCurrency)
	assert.Equal(t, "", result.Data[0].PreservedPricing)
	assert.Equal(t, "", result.Data[0].ProceedsReason)
	assert.Equal(t, "", result.Data[0].Client)
	assert.Equal(t, "iPhone", result.Data[0].Device)
	assert.Equal(t, " ", result.Data[0].State)
	assert.Equal(t, "RU", result.Data[0].Country)
	assert.Equal(t, 20, result.Data[0].ActiveStandardPriceSubscriptions.Value())
	assert.Equal(t, 0, result.Data[0].ActiveFreeTrialIntroductoryOfferSubscriptions.Value())
	assert.Equal(t, 0, result.Data[0].ActivePayUpFrontIntroductoryOfferSubscriptions.Value())
	assert.Equal(t, 0, result.Data[0].ActivePayAsYouGoIntroductoryOfferSubscriptions.Value())
	assert.Equal(t, 0, result.Data[0].FreeTrialPromotionalOfferSubscriptions.Value())
	assert.Equal(t, 0, result.Data[0].PayUpFrontPromotionalOfferSubscriptions.Value())
	assert.Equal(t, 0, result.Data[0].PayAsYouGoPromotionalOfferSubscriptions.Value())
	assert.Equal(t, 0, result.Data[0].MarketingOptIns.Value())
	assert.Equal(t, 0, result.Data[0].BillingRetry.Value())
	assert.Equal(t, 0, result.Data[0].GracePeriod.Value())

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(t, body)
}

func Test_Sales_GetSubscriptionsReports_Error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := buildStubConfig()
	transport := buildStubHttpTransport()
	ar := newResourceAbstract(transport, config)

	rsp := buildStubResponseFromFile(http.StatusBadRequest, "stubs/errors/invalid.parameter.json")
	rsp.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder("GET", config.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	resource := &SalesReportsResource{ar}
	filter := NewSubscriptionsReportsFilter()
	filter.SubTypeSummary().Version12().Daily()
	ctx := context.Background()
	result, resp, err := resource.GetSubscriptionsReports(ctx, filter)
	assert.Error(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	assert.Equal(t, "The version parameter you have specified is invalid. The latest version for this report is 1_0.", err.Error())

	assert.False(t, result.IsSuccess())
	assert.Equal(t, "The version parameter you have specified is invalid. The latest version for this report is 1_0.", result.GetError())
	assert.Len(t, result.Errors, 1)
	assert.Empty(t, result.Data)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(t, body)
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
