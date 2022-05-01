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
	assert.Equal(t, "APPLE", result.Data[0].Provider)
	assert.Equal(t, "US", result.Data[0].ProviderCountry)
	assert.Equal(t, "foo.bar.baz", result.Data[0].SKU)
	assert.Equal(t, " ", result.Data[0].Developer)
	assert.Equal(t, "FooBarTitle", result.Data[0].Title)
	assert.Equal(t, "", result.Data[0].Version)
	assert.Equal(t, "IAY", result.Data[0].ProductTypeIdentifier)
	assert.Equal(t, float64(12), result.Data[0].Units.Value())
	assert.Equal(t, 1234567890, result.Data[0].AppleIdentifier.Value())
	assert.Equal(t, 209.3000030517578, result.Data[0].DeveloperProceeds.Value())
	assert.Equal(t, "2020-10-05", result.Data[0].BeginDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, "2020-10-05", result.Data[0].EndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, "RUB", result.Data[0].CustomerCurrency)
	assert.Equal(t, "RU", result.Data[0].CountryCode)
	assert.Equal(t, "RUB", result.Data[0].CurrencyOfProceeds)
	assert.Equal(t, 1234567890, result.Data[0].AppleIdentifier.Value())
	assert.Equal(t, float64(299), result.Data[0].CustomerPrice.Value())
	assert.Equal(t, " ", result.Data[0].PromoCode)
	assert.Equal(t, "foo.bar.baz", result.Data[0].ParentIdentifier)
	assert.Equal(t, "Renewal", result.Data[0].Subscription)
	assert.Equal(t, "7 Days", result.Data[0].Period)
	assert.Equal(t, "Lifestyle", result.Data[0].Category)
	assert.Equal(t, "", result.Data[0].CMB)
	assert.Equal(t, "iPhone", result.Data[0].Device)
	assert.Equal(t, "iOS", result.Data[0].SupportedPlatforms)
	assert.Equal(t, " ", result.Data[0].ProceedsReason)
	assert.Equal(t, "Yes", result.Data[0].PreservedPricing)
	assert.Equal(t, " ", result.Data[0].Client)
	assert.Equal(t, " ", result.Data[0].OrderType)

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

func Test_Sales_GetSubscriptionsEventsReports_Success(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := buildStubConfig()
	transport := buildStubHttpTransport()
	ar := newResourceAbstract(transport, config)

	rsp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/subscriptions-events.tsv")
	rsp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", config.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	resource := &SalesReportsResource{ar}
	filter := NewSubscriptionsEventsReportsFilter()
	filter.SubTypeSummary().Version12().Daily()
	ctx := context.Background()
	result, resp, err := resource.GetSubscriptionsEventsReports(ctx, filter)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)

	assert.True(t, result.IsSuccess())
	assert.Empty(t, result.GetError())
	assert.Empty(t, result.Errors)
	assert.Equal(t, "2020-10-06", result.Data[0].EventDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, "Renew", result.Data[0].Event)
	assert.Equal(t, "AppFooBar", result.Data[0].AppName)
	assert.Equal(t, 1234567890, result.Data[0].AppAppleID.Value())
	assert.Equal(t, "foo.bar.baz", result.Data[0].SubscriptionName)
	assert.Equal(t, 1234567890, result.Data[0].SubscriptionAppleID.Value())
	assert.Equal(t, 1234567890, result.Data[0].SubscriptionGroupID.Value())
	assert.Equal(t, "7 Days", result.Data[0].StandardSubscriptionDuration)
	assert.Equal(t, " ", result.Data[0].PromotionalOfferName)
	assert.Equal(t, " ", result.Data[0].PromotionalOfferID)
	assert.Equal(t, "", result.Data[0].SubscriptionOfferType)
	assert.Equal(t, "", result.Data[0].SubscriptionOfferDuration)
	assert.Equal(t, "", result.Data[0].MarketingOptIn)
	assert.Equal(t, " ", result.Data[0].MarketingOptInDuration)
	assert.Equal(t, "", result.Data[0].PreservedPricing)
	assert.Equal(t, "", result.Data[0].ProceedsReason)
	assert.Equal(t, 11, result.Data[0].ConsecutivePaidPeriods.Value())
	assert.Equal(t, "2020-07-25", result.Data[0].OriginalStartDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, "", result.Data[0].Client)
	assert.Equal(t, "iPhone", result.Data[0].Device)
	assert.Equal(t, " ", result.Data[0].State)
	assert.Equal(t, "RU", result.Data[0].Country)
	assert.Equal(t, "", result.Data[0].PreviousSubscriptionName)
	assert.Equal(t, 0, result.Data[0].PreviousSubscriptionAppleID.Value())
	assert.Equal(t, 0, result.Data[0].DaysBeforeCanceling.Value())
	assert.Equal(t, " ", result.Data[0].CancellationReason)
	assert.Equal(t, 0, result.Data[0].DaysCanceled.Value())
	assert.Equal(t, 1, result.Data[0].Quantity.Value())

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(t, body)
}

func Test_Sales_GetSubscriptionsEventsReports_Error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := buildStubConfig()
	transport := buildStubHttpTransport()
	ar := newResourceAbstract(transport, config)

	rsp := buildStubResponseFromFile(http.StatusBadRequest, "stubs/errors/invalid.parameter.json")
	rsp.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder("GET", config.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	resource := &SalesReportsResource{ar}
	filter := NewSubscriptionsEventsReportsFilter()
	filter.SubTypeSummary().Version12().Daily()
	ctx := context.Background()
	result, resp, err := resource.GetSubscriptionsEventsReports(ctx, filter)
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

func Test_Sales_GetSubscribersReports_Success(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := buildStubConfig()
	transport := buildStubHttpTransport()
	ar := newResourceAbstract(transport, config)

	rsp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/subscribers.tsv")
	rsp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", config.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	resource := &SalesReportsResource{ar}
	filter := NewSubscribersReportsFilter()
	filter.SubTypeDetailed().Version12().Daily()
	ctx := context.Background()
	result, resp, err := resource.GetSubscribersReports(ctx, filter)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)

	assert.True(t, result.IsSuccess())
	assert.Empty(t, result.GetError())
	assert.Empty(t, result.Errors)
	assert.Equal(t, "2020-10-05", result.Data[0].EventDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, "FooBarApp", result.Data[0].AppName)
	assert.Equal(t, 1234567890, result.Data[0].AppAppleID.Value())
	assert.Equal(t, "foo.bar.baz", result.Data[0].SubscriptionName)
	assert.Equal(t, 1234567890, result.Data[0].SubscriptionAppleID.Value())
	assert.Equal(t, 1234567890, result.Data[0].SubscriptionGroupID.Value())
	assert.Equal(t, "7 Days", result.Data[0].StandardSubscriptionDuration)
	assert.Equal(t, "", result.Data[0].PromotionalOfferName)
	assert.Equal(t, "", result.Data[0].PromotionalOfferID)
	assert.Equal(t, "", result.Data[0].SubscriptionOfferType)
	assert.Equal(t, "", result.Data[0].SubscriptionOfferDuration)
	assert.Equal(t, "", result.Data[0].MarketingOptInDuration)
	assert.Equal(t, 4.489999771118164, result.Data[0].CustomerPrice.Value())
	assert.Equal(t, "USD", result.Data[0].CustomerCurrency)
	assert.Equal(t, 3.1500000953674316, result.Data[0].DeveloperProceeds.Value())
	assert.Equal(t, "USD", result.Data[0].ProceedsCurrency)
	assert.Equal(t, " ", result.Data[0].PreservedPricing)
	assert.Equal(t, " ", result.Data[0].ProceedsReason)
	assert.Equal(t, " ", result.Data[0].Client)
	assert.Equal(t, "UA", result.Data[0].Country)
	assert.Equal(t, 1234567890000, result.Data[0].SubscriberID.Value())
	assert.Equal(t, "", result.Data[0].SubscriberIDReset)
	assert.Equal(t, "", result.Data[0].Refund)
	//assert.Equal(t, "", result.Data[0].PurchaseDate.Value())
	assert.Equal(t, 1, result.Data[0].Units.Value())

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(t, body)
}

func Test_Sales_GetSubscribersReports_Error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := buildStubConfig()
	transport := buildStubHttpTransport()
	ar := newResourceAbstract(transport, config)

	rsp := buildStubResponseFromFile(http.StatusBadRequest, "stubs/errors/invalid.parameter.json")
	rsp.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder("GET", config.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	resource := &SalesReportsResource{ar}
	filter := NewSubscribersReportsFilter()
	filter.SubTypeDetailed().Version12().Daily()
	ctx := context.Background()
	result, resp, err := resource.GetSubscribersReports(ctx, filter)
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

func Test_Sales_GetPreOrdersReports_Success(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := buildStubConfig()
	transport := buildStubHttpTransport()
	ar := newResourceAbstract(transport, config)

	rsp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/preorders.tsv")
	rsp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", config.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	resource := &SalesReportsResource{ar}
	filter := NewPreOrdersReportsFilter()
	filter.SubTypeSummary().Version10().Daily()
	ctx := context.Background()
	result, resp, err := resource.GetPreOrdersReports(ctx, filter)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)

	assert.True(t, result.IsSuccess())
	assert.Empty(t, result.GetError())
	assert.Empty(t, result.Errors)
	assert.Equal(t, "APPLE", result.Data[0].Provider)
	assert.Equal(t, "RU", result.Data[0].ProviderCountry)
	assert.Equal(t, "Foo", result.Data[0].Title)
	assert.Equal(t, "", result.Data[0].SKU)
	assert.Equal(t, "", result.Data[0].Developer)
	assert.Equal(t, "2020-10-05", result.Data[0].PreOrderStartDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, "2020-10-05", result.Data[0].PreOrderEndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, 10.199999809265137, result.Data[0].Ordered.Value())
	assert.Equal(t, 5.5, result.Data[0].Canceled.Value())
	assert.Equal(t, float64(10), result.Data[0].CumulativeOrdered.Value())
	assert.Equal(t, float64(12), result.Data[0].CumulativeCanceled.Value())
	assert.Equal(t, "2020-10-05", result.Data[0].StartDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, "2020-10-05", result.Data[0].EndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, "RU", result.Data[0].CountryCode)
	assert.Equal(t, 1234567890, result.Data[0].AppleIdentifier.Value())
	assert.Equal(t, "Lifestyle", result.Data[0].Category)
	assert.Equal(t, "iPhone", result.Data[0].Device)
	assert.Equal(t, "iOS", result.Data[0].SupportedPlatforms)
	assert.Equal(t, "foo", result.Data[0].Client)
	assert.Equal(t, "RU", result.Data[0].ProviderCountry)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(t, body)
}

func Test_Sales_GetPreOrdersReports_Error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := buildStubConfig()
	transport := buildStubHttpTransport()
	ar := newResourceAbstract(transport, config)

	rsp := buildStubResponseFromFile(http.StatusBadRequest, "stubs/errors/invalid.parameter.json")
	rsp.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder("GET", config.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	resource := &SalesReportsResource{ar}
	filter := NewPreOrdersReportsFilter()
	filter.SubTypeSummary().Version10().Daily()
	ctx := context.Background()
	result, resp, err := resource.GetPreOrdersReports(ctx, filter)
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
