package appstore

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type SalesReportsBaseFilterSuite struct {
	suite.Suite
	testable *SalesReportsBaseFilter
}

func (suite *SalesReportsBaseFilterSuite) SetupTest() {
	suite.testable = &SalesReportsBaseFilter{}
}

func (suite *SalesReportsBaseFilterSuite) TestFillByDefault() {
	suite.testable.Daily().TypeSales().SubTypeSummary().Version10()
	assert.Equal(suite.T(), SalesReportTypeSales, suite.testable.ReportType)
	assert.Equal(suite.T(), SalesReportSubTypeSummary, suite.testable.ReportSubType)
	assert.Equal(suite.T(), SalesReportFrequencyDaily, suite.testable.Frequency)
	assert.Equal(suite.T(), SalesReportVersion10, suite.testable.Version)
}

func (suite *SalesReportsBaseFilterSuite) TestToQueryParamsMapOnlyRequired() {
	suite.testable.Yearly().TypeSales().SubTypeSummary()
	qs := make(map[string]interface{})
	qs["filter[reportSubType]"] = string(SalesReportSubTypeSummary)
	qs["filter[reportType]"] = string(SalesReportTypeSales)
	qs["filter[frequency]"] = string(SalesReportFrequencyYearly)
	assert.Equal(suite.T(), qs, suite.testable.ToQueryParamsMap())
}

func (suite *SalesReportsBaseFilterSuite) TestToQueryParamsMap() {
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	suite.testable.Daily().TypeSales().SubTypeSummary().Version10().SetReportDate(date)
	qs := make(map[string]interface{})
	qs["filter[reportDate]"] = "2020-05-05"
	qs["filter[reportSubType]"] = string(SalesReportSubTypeSummary)
	qs["filter[reportType]"] = string(SalesReportTypeSales)
	qs["filter[frequency]"] = string(SalesReportFrequencyDaily)
	qs["filter[version]"] = string(SalesReportVersion10)
	assert.Equal(suite.T(), qs, suite.testable.ToQueryParamsMap())
}

func (suite *SalesReportsBaseFilterSuite) TestSetSubType() {
	suite.testable.SubTypeSummary()
	assert.Equal(suite.T(), suite.testable.ReportSubType, SalesReportSubTypeSummary)
	suite.testable.SubTypeDetailed()
	assert.Equal(suite.T(), suite.testable.ReportSubType, SalesReportSubTypeDetailed)
	suite.testable.SubTypeOptIn()
	assert.Equal(suite.T(), suite.testable.ReportSubType, SalesReportSubTypeOptIn)
}

func (suite *SalesReportsBaseFilterSuite) TestSetType() {
	suite.testable.TypeSales()
	assert.Equal(suite.T(), suite.testable.ReportType, SalesReportTypeSales)
	suite.testable.TypeNewsStand()
	assert.Equal(suite.T(), suite.testable.ReportType, SalesReportTypeNewsStand)
	suite.testable.TypePreOrder()
	assert.Equal(suite.T(), suite.testable.ReportType, SalesReportTypePreorder)
	suite.testable.TypeSubscriber()
	assert.Equal(suite.T(), suite.testable.ReportType, SalesReportTypeSubscriber)
	suite.testable.TypeSubscription()
	assert.Equal(suite.T(), suite.testable.ReportType, SalesReportTypeSubscription)
	suite.testable.TypeSubscriptionEvent()
	assert.Equal(suite.T(), suite.testable.ReportType, SalesReportTypeSubscriptionEvent)
}

func (suite *SalesReportsBaseFilterSuite) TestSetFrequency() {
	suite.testable.Daily()
	assert.Equal(suite.T(), suite.testable.Frequency, SalesReportFrequencyDaily)
	suite.testable.Weekly()
	assert.Equal(suite.T(), suite.testable.Frequency, SalesReportFrequencyWeekly)
	suite.testable.Monthly()
	assert.Equal(suite.T(), suite.testable.Frequency, SalesReportFrequencyMonthly)
	suite.testable.Yearly()
	assert.Equal(suite.T(), suite.testable.Frequency, SalesReportFrequencyYearly)
}

func (suite *SalesReportsBaseFilterSuite) TestSetVersion() {
	suite.testable.Version10()
	assert.Equal(suite.T(), suite.testable.Version, SalesReportVersion10)
	suite.testable.Version12()
	assert.Equal(suite.T(), suite.testable.Version, SalesReportVersion12)
	suite.testable.Version13()
	assert.Equal(suite.T(), suite.testable.Version, SalesReportVersion13)
}

func (suite *SalesReportsBaseFilterSuite) TestIsValid() {
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	suite.testable.Daily().TypeSales().SubTypeSummary().Version10().SetReportDate(date)
	err := suite.testable.IsValid()
	assert.Nil(suite.T(), err)
}

func (suite *SalesReportsBaseFilterSuite) TestIsInvalidReportType() {
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	suite.testable.Daily().SubTypeSummary().Version10().SetReportDate(date)
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SalesReportsBaseFilter.IsValid: ReportType is required", err.Error())
}

func (suite *SalesReportsBaseFilterSuite) TestIsInvalidReportSubType() {
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	suite.testable.Daily().TypeSales().Version10().SetReportDate(date)
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SalesReportsBaseFilter.IsValid: ReportSubType is required", err.Error())
}

func (suite *SalesReportsBaseFilterSuite) TestIsInvalidFrequency() {
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	suite.testable.SubTypeSummary().TypeSales().Version10().SetReportDate(date)
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SalesReportsBaseFilter.IsValid: Frequency is required", err.Error())
}

func TestSalesReportsBaseFilterSuite(t *testing.T) {
	suite.Run(t, new(SalesReportsBaseFilterSuite))
}

func Test_Sales_SalesReportsFilter_IsInvalidReportType(t *testing.T) {
	filter := NewSalesReportsFilter()
	filter.SubTypeSummary().Version10().Daily()
	filter.TypePreOrder()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SalesReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func Test_Sales_SalesReportsFilter_IsInvalidReportSubType(t *testing.T) {
	filter := NewSalesReportsFilter()
	filter.Version10().Daily()
	filter.SubTypeDetailed()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SalesReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func Test_Sales_SalesReportsFilter_IsInvalidVersion(t *testing.T) {
	filter := NewSalesReportsFilter()
	filter.SubTypeSummary().Daily()
	filter.Version12()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SalesReportsFilter.IsValid: Version is not valid", err.Error())
}

func Test_Sales_SubscriptionsReportsFilter_IsInvalidReportType(t *testing.T) {
	filter := NewSubscriptionsReportsFilter()
	filter.SubTypeSummary().Version12().Daily()
	filter.TypePreOrder()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscriptionsReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func Test_Sales_SubscriptionsReportsFilter_IsInvalidReportSubType(t *testing.T) {
	filter := NewSubscriptionsReportsFilter()
	filter.Version12().Daily()
	filter.SubTypeOptIn()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscriptionsReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func Test_Sales_SubscriptionsReportsFilter_IsInvalidFrequency(t *testing.T) {
	filter := NewSubscriptionsReportsFilter()
	filter.SubTypeSummary().Version12()
	filter.Yearly()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscriptionsReportsFilter.IsValid: Frequency is not valid", err.Error())
}

func Test_Sales_SubscriptionsReportsFilter_IsInvalidVersion(t *testing.T) {
	filter := NewSubscriptionsReportsFilter()
	filter.SubTypeSummary().Daily()
	filter.Version10()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscriptionsReportsFilter.IsValid: Version is not valid", err.Error())
}

func Test_Sales_SubscriptionsEventsReportsFilter_IsInvalidReportType(t *testing.T) {
	filter := NewSubscriptionsEventsReportsFilter()
	filter.SubTypeSummary().Version12().Daily()
	filter.TypePreOrder()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscriptionsEventsReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func Test_Sales_SubscriptionsEventsReportsFilter_IsInvalidReportSubType(t *testing.T) {
	filter := NewSubscriptionsEventsReportsFilter()
	filter.Version12().Daily()
	filter.SubTypeOptIn()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscriptionsEventsReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func Test_Sales_SubscriptionsEventsReportsFilter_IsInvalidReportFrequency(t *testing.T) {
	filter := NewSubscriptionsEventsReportsFilter()
	filter.SubTypeSummary().Version12()
	filter.Yearly()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscriptionsEventsReportsFilter.IsValid: Frequency is not valid", err.Error())
}

func Test_Sales_SubscriptionsEventsReportsFilter_IsInvalidReportVersion(t *testing.T) {
	filter := NewSubscriptionsEventsReportsFilter()
	filter.SubTypeSummary().Daily()
	filter.Version10()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscriptionsEventsReportsFilter.IsValid: Version is not valid", err.Error())
}

func Test_Sales_SubscribersReportsFilter_IsInvalidReportType(t *testing.T) {
	filter := NewSubscribersReportsFilter()
	filter.SubTypeDetailed().Version12().Daily()
	filter.TypePreOrder()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscribersReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func Test_Sales_SubscribersReportsFilter_IsInvalidReportSubType(t *testing.T) {
	filter := NewSubscribersReportsFilter()
	filter.Version12().Daily()
	filter.SubTypeSummary()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscribersReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func Test_Sales_SubscribersReportsFilter_IsInvalidReportFrequency(t *testing.T) {
	filter := NewSubscribersReportsFilter()
	filter.SubTypeDetailed().Version12()
	filter.Yearly()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscribersReportsFilter.IsValid: Frequency is not valid", err.Error())
}

func Test_Sales_SubscribersReportsFilter_IsInvalidReportVersion(t *testing.T) {
	filter := NewSubscribersReportsFilter()
	filter.SubTypeDetailed().Daily()
	filter.Version10()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscribersReportsFilter.IsValid: Version is not valid", err.Error())
}

func Test_Sales_SubscriptionsOffersCodesRedemptionReportsFilter_IsInvalidReportType(t *testing.T) {
	filter := NewSubscriptionsOffersCodesRedemptionReportsFilter()
	filter.SubTypeDetailed().Version12().Daily()
	filter.TypePreOrder()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func Test_Sales_SubscriptionsOffersCodesRedemptionReportsFilter_IsInvalidReportSubType(t *testing.T) {
	filter := NewSubscriptionsOffersCodesRedemptionReportsFilter()
	filter.Version10().Daily()
	filter.SubTypeDetailed()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func Test_Sales_SubscriptionsOffersCodesRedemptionReportsFilterr_IsInvalidReportFrequency(t *testing.T) {
	filter := NewSubscriptionsOffersCodesRedemptionReportsFilter()
	filter.SubTypeSummary().Version10()
	filter.Yearly()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: Frequency is not valid", err.Error())
}

func Test_Sales_SubscriptionsOffersCodesRedemptionReportsFilter_IsInvalidReportVersion(t *testing.T) {
	filter := NewSubscriptionsOffersCodesRedemptionReportsFilter()
	filter.SubTypeSummary().Daily()
	filter.Version12()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: Version is not valid", err.Error())
}

func Test_Sales_NewsstandReportsFilter_IsInvalidReportType(t *testing.T) {
	filter := NewNewsstandReportsFilter()
	filter.SubTypeDetailed().Version10().Daily()
	filter.TypePreOrder()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "NewsstandReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func Test_Sales_NewsstandReportsFilter_IsInvalidReportSubType(t *testing.T) {
	filter := NewNewsstandReportsFilter()
	filter.Version10().Daily()
	filter.SubTypeSummary()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "NewsstandReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func Test_Sales_NewsstandReportsFilter_IsInvalidReportFrequency(t *testing.T) {
	filter := NewNewsstandReportsFilter()
	filter.SubTypeDetailed().Version10()
	filter.Yearly()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "NewsstandReportsFilter.IsValid: Frequency is not valid", err.Error())
}

func Test_Sales_NewsstandReportsFilter_IsInvalidReportVersion(t *testing.T) {
	filter := NewNewsstandReportsFilter()
	filter.SubTypeDetailed().Daily()
	filter.Version12()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "NewsstandReportsFilter.IsValid: Version is not valid", err.Error())
}

func Test_Sales_PreOrdersReportsFilter_IsInvalidReportType(t *testing.T) {
	filter := NewPreOrdersReportsFilter()
	filter.SubTypeSummary().Version10().Daily()
	filter.TypeNewsStand()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "PreOrdersReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func Test_Sales_PreOrdersReportsFilter_IsInvalidReportSubType(t *testing.T) {
	filter := NewPreOrdersReportsFilter()
	filter.Version10().Daily()
	filter.SubTypeDetailed()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "PreOrdersReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func Test_Sales_PreOrdersReportsFilter_IsInvalidReportVersion(t *testing.T) {
	filter := NewPreOrdersReportsFilter()
	filter.SubTypeSummary().Daily()
	filter.Version12()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "PreOrdersReportsFilter.IsValid: Version is not valid", err.Error())
}
