package appstore

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
)

type SalesReportsResourceTestSuite struct {
	suite.Suite
	cfg      *Config
	ctx      context.Context
	testable *SalesReportsResource
}

func (suite *SalesReportsResourceTestSuite) SetupTest() {
	suite.cfg = buildStubConfig()
	suite.ctx = context.Background()
	suite.testable = buildStubSalesReportsResource()
	httpmock.Activate()
}

func (suite *SalesReportsResourceTestSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (suite *SalesReportsResourceTestSuite) TestGetReportsInvalidFilter() {
	filter := &SalesReportsBaseFilter{}
	filter.TypeSales().SubTypeSummary().Version10()
	resp, err := suite.testable.GetReports(suite.ctx, filter)
	assert.Nil(suite.T(), resp)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SalesReportsResource.GetReports invalid filter: SalesReportsBaseFilter.IsValid: Frequency is required", err.Error())
}

func (suite *SalesReportsResourceTestSuite) TestGetReportsSuccess() {
	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/sales.tsv")
	resp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(resp))

	filter := &SalesReportsBaseFilter{}
	filter.TypeSales().SubTypeSummary().Version10().Daily()
	rsp, err := suite.testable.GetReports(suite.ctx, filter)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), rsp)
}

func (suite *SalesReportsResourceTestSuite) TestGetSalesReportsSuccess() {
	rsp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/sales.tsv")
	rsp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	filter := NewSalesReportsFilter()
	filter.SubTypeSummary().Version11().Daily()

	result, resp, err := suite.testable.GetSalesReports(suite.ctx, filter)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)

	assert.True(suite.T(), result.IsSuccess())
	assert.Empty(suite.T(), result.GetError())
	assert.Empty(suite.T(), result.Errors)
	assert.Equal(suite.T(), "APPLE", result.Data[0].Provider)
	assert.Equal(suite.T(), "US", result.Data[0].ProviderCountry)
	assert.Equal(suite.T(), "foo.bar.baz", result.Data[0].SKU)
	assert.Equal(suite.T(), " ", result.Data[0].Developer)
	assert.Equal(suite.T(), "FooBarTitle", result.Data[0].Title)
	assert.Equal(suite.T(), "", result.Data[0].Version)
	assert.Equal(suite.T(), "IAY", result.Data[0].ProductTypeIdentifier)
	assert.Equal(suite.T(), float64(12), result.Data[0].Units.Value())
	assert.Equal(suite.T(), 1234567890, result.Data[0].AppleIdentifier.Value())
	assert.Equal(suite.T(), 209.3000030517578, result.Data[0].DeveloperProceeds.Value())
	assert.Equal(suite.T(), "2020-10-05", result.Data[0].BeginDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), "2020-10-05", result.Data[0].EndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), "RUB", result.Data[0].CustomerCurrency)
	assert.Equal(suite.T(), "RU", result.Data[0].CountryCode)
	assert.Equal(suite.T(), "RUB", result.Data[0].CurrencyOfProceeds)
	assert.Equal(suite.T(), 1234567890, result.Data[0].AppleIdentifier.Value())
	assert.Equal(suite.T(), float64(299), result.Data[0].CustomerPrice.Value())
	assert.Equal(suite.T(), " ", result.Data[0].PromoCode)
	assert.Equal(suite.T(), "foo.bar.baz", result.Data[0].ParentIdentifier)
	assert.Equal(suite.T(), "Renewal", result.Data[0].Subscription)
	assert.Equal(suite.T(), "7 Days", result.Data[0].Period)
	assert.Equal(suite.T(), "Lifestyle", result.Data[0].Category)
	assert.Equal(suite.T(), "", result.Data[0].CMB)
	assert.Equal(suite.T(), "iPhone", result.Data[0].Device)
	assert.Equal(suite.T(), "iOS", result.Data[0].SupportedPlatforms)
	assert.Equal(suite.T(), " ", result.Data[0].ProceedsReason)
	assert.Equal(suite.T(), "Yes", result.Data[0].PreservedPricing)
	assert.Equal(suite.T(), " ", result.Data[0].Client)
	assert.Equal(suite.T(), " ", result.Data[0].OrderType)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
}

func (suite *SalesReportsResourceTestSuite) TestGetSalesReportsError() {
	rsp := buildStubResponseFromFile(http.StatusBadRequest, "stubs/errors/invalid.parameter.json")
	rsp.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	filter := NewSalesReportsFilter()
	filter.SubTypeSummary().Version11().Daily()

	result, resp, err := suite.testable.GetSalesReports(suite.ctx, filter)
	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), "The version parameter you have specified is invalid. The latest version for this report is 1_0.", err.Error())

	assert.False(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "The version parameter you have specified is invalid. The latest version for this report is 1_0.", result.GetError())
	assert.Len(suite.T(), result.Errors, 1)
	assert.Empty(suite.T(), result.Data)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
}

func (suite *SalesReportsResourceTestSuite) TestGetSalesReportsInvalidFilter() {
	filter := NewSalesReportsFilter()
	filter.Version10().Daily()
	result, resp, err := suite.testable.GetSalesReports(suite.ctx, filter)
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), resp)
	assert.Empty(suite.T(), result)
	assert.Equal(suite.T(), "SalesReportsResource.GetSalesReports error: SalesReportsResource.GetReports invalid filter: SalesReportsBaseFilter.IsValid: ReportSubType is required", err.Error())
}

func (suite *SalesReportsResourceTestSuite) TestGetSubscriptionsReportsSuccess() {
	rsp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/subscriptions.tsv")
	rsp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	filter := NewSubscriptionsReportsFilter()
	filter.SubTypeSummary().Version12().Daily()

	result, resp, err := suite.testable.GetSubscriptionsReports(suite.ctx, filter)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)

	assert.True(suite.T(), result.IsSuccess())
	assert.Empty(suite.T(), result.GetError())
	assert.Empty(suite.T(), result.Errors)
	assert.Equal(suite.T(), "FooBarApp", result.Data[0].AppName)
	assert.Equal(suite.T(), 1234567890, result.Data[0].AppAppleID.Value())
	assert.Equal(suite.T(), "foo.bar.baz", result.Data[0].SubscriptionName)
	assert.Equal(suite.T(), 1234567890, result.Data[0].SubscriptionAppleID.Value())
	assert.Equal(suite.T(), 1234567890, result.Data[0].SubscriptionGroupID.Value())
	assert.Equal(suite.T(), "1 Year", result.Data[0].StandardSubscriptionDuration)
	assert.Equal(suite.T(), " ", result.Data[0].PromotionalOfferName)
	assert.Equal(suite.T(), " ", result.Data[0].PromotionalOfferID)
	assert.Equal(suite.T(), float64(2950), result.Data[0].CustomerPrice.Value())
	assert.Equal(suite.T(), "RUB", result.Data[0].CustomerCurrency)
	assert.Equal(suite.T(), 2065.00, result.Data[0].DeveloperProceeds.Value())
	assert.Equal(suite.T(), "RUB", result.Data[0].ProceedsCurrency)
	assert.Equal(suite.T(), "", result.Data[0].PreservedPricing)
	assert.Equal(suite.T(), "", result.Data[0].ProceedsReason)
	assert.Equal(suite.T(), "", result.Data[0].Client)
	assert.Equal(suite.T(), "iPhone", result.Data[0].Device)
	assert.Equal(suite.T(), " ", result.Data[0].State)
	assert.Equal(suite.T(), "RU", result.Data[0].Country)
	assert.Equal(suite.T(), 20, result.Data[0].ActiveStandardPriceSubscriptions.Value())
	assert.Equal(suite.T(), 0, result.Data[0].ActiveFreeTrialIntroductoryOfferSubscriptions.Value())
	assert.Equal(suite.T(), 0, result.Data[0].ActivePayUpFrontIntroductoryOfferSubscriptions.Value())
	assert.Equal(suite.T(), 0, result.Data[0].ActivePayAsYouGoIntroductoryOfferSubscriptions.Value())
	assert.Equal(suite.T(), 0, result.Data[0].FreeTrialPromotionalOfferSubscriptions.Value())
	assert.Equal(suite.T(), 0, result.Data[0].PayUpFrontPromotionalOfferSubscriptions.Value())
	assert.Equal(suite.T(), 0, result.Data[0].PayAsYouGoPromotionalOfferSubscriptions.Value())
	assert.Equal(suite.T(), 0, result.Data[0].MarketingOptIns.Value())
	assert.Equal(suite.T(), 0, result.Data[0].BillingRetry.Value())
	assert.Equal(suite.T(), 0, result.Data[0].GracePeriod.Value())

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
}

func (suite *SalesReportsResourceTestSuite) TestGetSubscriptionsReportsError() {
	rsp := buildStubResponseFromFile(http.StatusBadRequest, "stubs/errors/invalid.parameter.json")
	rsp.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	filter := NewSubscriptionsReportsFilter()
	filter.SubTypeSummary().Version12().Daily()
	result, resp, err := suite.testable.GetSubscriptionsReports(suite.ctx, filter)
	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), "The version parameter you have specified is invalid. The latest version for this report is 1_0.", err.Error())

	assert.False(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "The version parameter you have specified is invalid. The latest version for this report is 1_0.", result.GetError())
	assert.Len(suite.T(), result.Errors, 1)
	assert.Empty(suite.T(), result.Data)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
}

func (suite *SalesReportsResourceTestSuite) TestGetSubscriptionsEventsReportsSuccess() {
	rsp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/subscriptions-events.tsv")
	rsp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	filter := NewSubscriptionsEventsReportsFilter()
	filter.SubTypeSummary().Version12().Daily()

	result, resp, err := suite.testable.GetSubscriptionsEventsReports(suite.ctx, filter)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)

	assert.True(suite.T(), result.IsSuccess())
	assert.Empty(suite.T(), result.GetError())
	assert.Empty(suite.T(), result.Errors)
	assert.Equal(suite.T(), "2020-10-06", result.Data[0].EventDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), "Renew", result.Data[0].Event)
	assert.Equal(suite.T(), "AppFooBar", result.Data[0].AppName)
	assert.Equal(suite.T(), 1234567890, result.Data[0].AppAppleID.Value())
	assert.Equal(suite.T(), "foo.bar.baz", result.Data[0].SubscriptionName)
	assert.Equal(suite.T(), 1234567890, result.Data[0].SubscriptionAppleID.Value())
	assert.Equal(suite.T(), 1234567890, result.Data[0].SubscriptionGroupID.Value())
	assert.Equal(suite.T(), "7 Days", result.Data[0].StandardSubscriptionDuration)
	assert.Equal(suite.T(), " ", result.Data[0].PromotionalOfferName)
	assert.Equal(suite.T(), " ", result.Data[0].PromotionalOfferID)
	assert.Equal(suite.T(), "", result.Data[0].SubscriptionOfferType)
	assert.Equal(suite.T(), "", result.Data[0].SubscriptionOfferDuration)
	assert.Equal(suite.T(), "", result.Data[0].MarketingOptIn)
	assert.Equal(suite.T(), " ", result.Data[0].MarketingOptInDuration)
	assert.Equal(suite.T(), "", result.Data[0].PreservedPricing)
	assert.Equal(suite.T(), "", result.Data[0].ProceedsReason)
	assert.Equal(suite.T(), 11, result.Data[0].ConsecutivePaidPeriods.Value())
	assert.Equal(suite.T(), "2020-07-25", result.Data[0].OriginalStartDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), "", result.Data[0].Client)
	assert.Equal(suite.T(), "iPhone", result.Data[0].Device)
	assert.Equal(suite.T(), " ", result.Data[0].State)
	assert.Equal(suite.T(), "RU", result.Data[0].Country)
	assert.Equal(suite.T(), "", result.Data[0].PreviousSubscriptionName)
	assert.Equal(suite.T(), 0, result.Data[0].PreviousSubscriptionAppleID.Value())
	assert.Equal(suite.T(), 0, result.Data[0].DaysBeforeCanceling.Value())
	assert.Equal(suite.T(), " ", result.Data[0].CancellationReason)
	assert.Equal(suite.T(), 0, result.Data[0].DaysCanceled.Value())
	assert.Equal(suite.T(), 1, result.Data[0].Quantity.Value())

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
}

func (suite *SalesReportsResourceTestSuite) TestGetSubscriptionsEventsReportsError() {
	rsp := buildStubResponseFromFile(http.StatusBadRequest, "stubs/errors/invalid.parameter.json")
	rsp.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	filter := NewSubscriptionsEventsReportsFilter()
	filter.SubTypeSummary().Version12().Daily()

	result, resp, err := suite.testable.GetSubscriptionsEventsReports(suite.ctx, filter)
	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), "The version parameter you have specified is invalid. The latest version for this report is 1_0.", err.Error())

	assert.False(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "The version parameter you have specified is invalid. The latest version for this report is 1_0.", result.GetError())
	assert.Len(suite.T(), result.Errors, 1)
	assert.Empty(suite.T(), result.Data)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
}

func (suite *SalesReportsResourceTestSuite) TestGetSubscribersReportsSuccess() {
	rsp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/subscribers.tsv")
	rsp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	filter := NewSubscribersReportsFilter()
	filter.SubTypeDetailed().Version12().Daily()

	result, resp, err := suite.testable.GetSubscribersReports(suite.ctx, filter)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)

	assert.True(suite.T(), result.IsSuccess())
	assert.Empty(suite.T(), result.GetError())
	assert.Empty(suite.T(), result.Errors)
	assert.Equal(suite.T(), "2020-10-05", result.Data[0].EventDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), "FooBarApp", result.Data[0].AppName)
	assert.Equal(suite.T(), 1234567890, result.Data[0].AppAppleID.Value())
	assert.Equal(suite.T(), "foo.bar.baz", result.Data[0].SubscriptionName)
	assert.Equal(suite.T(), 1234567890, result.Data[0].SubscriptionAppleID.Value())
	assert.Equal(suite.T(), 1234567890, result.Data[0].SubscriptionGroupID.Value())
	assert.Equal(suite.T(), "7 Days", result.Data[0].StandardSubscriptionDuration)
	assert.Equal(suite.T(), "", result.Data[0].PromotionalOfferName)
	assert.Equal(suite.T(), "", result.Data[0].PromotionalOfferID)
	assert.Equal(suite.T(), "", result.Data[0].SubscriptionOfferType)
	assert.Equal(suite.T(), "", result.Data[0].SubscriptionOfferDuration)
	assert.Equal(suite.T(), "", result.Data[0].MarketingOptInDuration)
	assert.Equal(suite.T(), 4.489999771118164, result.Data[0].CustomerPrice.Value())
	assert.Equal(suite.T(), "USD", result.Data[0].CustomerCurrency)
	assert.Equal(suite.T(), 3.1500000953674316, result.Data[0].DeveloperProceeds.Value())
	assert.Equal(suite.T(), "USD", result.Data[0].ProceedsCurrency)
	assert.Equal(suite.T(), " ", result.Data[0].PreservedPricing)
	assert.Equal(suite.T(), " ", result.Data[0].ProceedsReason)
	assert.Equal(suite.T(), " ", result.Data[0].Client)
	assert.Equal(suite.T(), "UA", result.Data[0].Country)
	assert.Equal(suite.T(), 1234567890000, result.Data[0].SubscriberID.Value())
	assert.Equal(suite.T(), "", result.Data[0].SubscriberIDReset)
	assert.Equal(suite.T(), "", result.Data[0].Refund)
	//assert.Equal(t, "", result.Data[0].PurchaseDate.Value())
	assert.Equal(suite.T(), 1, result.Data[0].Units.Value())

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
}

func (suite *SalesReportsResourceTestSuite) TestGetSubscribersReportsError() {
	rsp := buildStubResponseFromFile(http.StatusBadRequest, "stubs/errors/invalid.parameter.json")
	rsp.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	filter := NewSubscribersReportsFilter()
	filter.SubTypeDetailed().Version12().Daily()

	result, resp, err := suite.testable.GetSubscribersReports(suite.ctx, filter)
	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), "The version parameter you have specified is invalid. The latest version for this report is 1_0.", err.Error())

	assert.False(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "The version parameter you have specified is invalid. The latest version for this report is 1_0.", result.GetError())
	assert.Len(suite.T(), result.Errors, 1)
	assert.Empty(suite.T(), result.Data)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
}

func (suite *SalesReportsResourceTestSuite) TestGetPreOrdersReportsSuccess() {
	rsp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/sales/preorders.tsv")
	rsp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	filter := NewPreOrdersReportsFilter()
	filter.SubTypeSummary().Version10().Daily()
	result, resp, err := suite.testable.GetPreOrdersReports(suite.ctx, filter)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)

	assert.True(suite.T(), result.IsSuccess())
	assert.Empty(suite.T(), result.GetError())
	assert.Empty(suite.T(), result.Errors)
	assert.Equal(suite.T(), "APPLE", result.Data[0].Provider)
	assert.Equal(suite.T(), "RU", result.Data[0].ProviderCountry)
	assert.Equal(suite.T(), "Foo", result.Data[0].Title)
	assert.Equal(suite.T(), "", result.Data[0].SKU)
	assert.Equal(suite.T(), "", result.Data[0].Developer)
	assert.Equal(suite.T(), "2020-10-05", result.Data[0].PreOrderStartDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), "2020-10-05", result.Data[0].PreOrderEndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), 10.199999809265137, result.Data[0].Ordered.Value())
	assert.Equal(suite.T(), 5.5, result.Data[0].Canceled.Value())
	assert.Equal(suite.T(), float64(10), result.Data[0].CumulativeOrdered.Value())
	assert.Equal(suite.T(), float64(12), result.Data[0].CumulativeCanceled.Value())
	assert.Equal(suite.T(), "2020-10-05", result.Data[0].StartDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), "2020-10-05", result.Data[0].EndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), "RU", result.Data[0].CountryCode)
	assert.Equal(suite.T(), 1234567890, result.Data[0].AppleIdentifier.Value())
	assert.Equal(suite.T(), "Lifestyle", result.Data[0].Category)
	assert.Equal(suite.T(), "iPhone", result.Data[0].Device)
	assert.Equal(suite.T(), "iOS", result.Data[0].SupportedPlatforms)
	assert.Equal(suite.T(), "foo", result.Data[0].Client)
	assert.Equal(suite.T(), "RU", result.Data[0].ProviderCountry)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
}

func (suite *SalesReportsResourceTestSuite) TestGetPreOrdersReportsError() {
	rsp := buildStubResponseFromFile(http.StatusBadRequest, "stubs/errors/invalid.parameter.json")
	rsp.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/v1/salesReports", httpmock.ResponderFromResponse(rsp))

	filter := NewPreOrdersReportsFilter()
	filter.SubTypeSummary().Version10().Daily()

	result, resp, err := suite.testable.GetPreOrdersReports(suite.ctx, filter)
	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), "The version parameter you have specified is invalid. The latest version for this report is 1_0.", err.Error())

	assert.False(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "The version parameter you have specified is invalid. The latest version for this report is 1_0.", result.GetError())
	assert.Len(suite.T(), result.Errors, 1)
	assert.Empty(suite.T(), result.Data)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
}

func (suite *SalesReportsResourceTestSuite) TestBuildQueryParams() {
	filter := &SalesReportsBaseFilter{}
	filter.TypeSales().SubTypeSummary().Version10().Daily()
	result := suite.testable.buildQueryParams(filter)
	qs := make(map[string]interface{})
	qs["filter[reportSubType]"] = string(SalesReportSubTypeSummary)
	qs["filter[reportType]"] = string(SalesReportTypeSales)
	qs["filter[frequency]"] = string(SalesReportFrequencyDaily)
	qs["filter[version]"] = string(SalesReportVersion10)
	qs["filter[vendorNumber]"] = suite.cfg.VendorNo
	assert.Equal(suite.T(), qs, result)
}

func TestSalesReportsResourceTestSuite(t *testing.T) {
	suite.Run(t, new(SalesReportsResourceTestSuite))
}
