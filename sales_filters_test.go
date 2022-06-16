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

type SalesReportsFilterSuite struct {
	suite.Suite
	testable *SalesReportsFilter
}

func (suite *SalesReportsFilterSuite) SetupTest() {
	suite.testable = NewSalesReportsFilter()
}

func (suite *SalesReportsFilterSuite) TestIsInvalidReportType() {
	suite.testable.SubTypeSummary().Version10().Daily()
	suite.testable.TypePreOrder()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SalesReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func (suite *SalesReportsFilterSuite) TestIsInvalidReportSubType() {
	suite.testable.Version10().Daily()
	suite.testable.SubTypeDetailed()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SalesReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func (suite *SalesReportsFilterSuite) TestIsInvalidVersion() {
	suite.testable.SubTypeSummary().Daily()
	suite.testable.Version12()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SalesReportsFilter.IsValid: Version is not valid", err.Error())
}

func TestSalesReportsFilterSuite(t *testing.T) {
	suite.Run(t, new(SalesReportsFilterSuite))
}

type SalesSubscriptionsReportsFilterSuite struct {
	suite.Suite
	testable *SubscriptionsReportsFilter
}

func (suite *SalesSubscriptionsReportsFilterSuite) SetupTest() {
	suite.testable = NewSubscriptionsReportsFilter()
}

func (suite *SalesSubscriptionsReportsFilterSuite) TestIsInvalidReportType() {
	suite.testable.SubTypeSummary().Version12().Daily()
	suite.testable.TypePreOrder()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func (suite *SalesSubscriptionsReportsFilterSuite) TestIsInvalidReportSubType() {
	suite.testable.Version12().Daily()
	suite.testable.SubTypeOptIn()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func (suite *SalesSubscriptionsReportsFilterSuite) TestIsInvalidFrequency() {
	suite.testable.SubTypeSummary().Version12()
	suite.testable.Yearly()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsReportsFilter.IsValid: Frequency is not valid", err.Error())
}

func (suite *SalesSubscriptionsReportsFilterSuite) TestIsInvalidVersion() {
	suite.testable.SubTypeSummary().Daily()
	suite.testable.Version10()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsReportsFilter.IsValid: Version is not valid", err.Error())
}

func TestSalesSubscriptionsReportsFilterSuite(t *testing.T) {
	suite.Run(t, new(SalesSubscriptionsReportsFilterSuite))
}

type SalesSubscriptionsEventsReportsFilterSuite struct {
	suite.Suite
	testable *SubscriptionsEventsReportsFilter
}

func (suite *SalesSubscriptionsEventsReportsFilterSuite) SetupTest() {
	suite.testable = NewSubscriptionsEventsReportsFilter()
}

func (suite *SalesSubscriptionsEventsReportsFilterSuite) TestIsInvalidReportType() {
	suite.testable.SubTypeSummary().Version12().Daily()
	suite.testable.TypePreOrder()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsEventsReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func (suite *SalesSubscriptionsEventsReportsFilterSuite) TestIsInvalidReportSubType() {
	suite.testable.Version12().Daily()
	suite.testable.SubTypeOptIn()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsEventsReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func (suite *SalesSubscriptionsEventsReportsFilterSuite) TestIsInvalidReportFrequency() {
	suite.testable.SubTypeSummary().Version12()
	suite.testable.Yearly()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsEventsReportsFilter.IsValid: Frequency is not valid", err.Error())
}

func (suite *SalesSubscriptionsEventsReportsFilterSuite) TestIsInvalidReportVersion() {
	suite.testable.SubTypeSummary().Daily()
	suite.testable.Version10()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsEventsReportsFilter.IsValid: Version is not valid", err.Error())
}

func TestSalesSubscriptionsEventsReportsFilterSuite(t *testing.T) {
	suite.Run(t, new(SalesSubscriptionsEventsReportsFilterSuite))
}

type SalesSubscribersReportsFilterSuite struct {
	suite.Suite
	testable *SubscribersReportsFilter
}

func (suite *SalesSubscribersReportsFilterSuite) SetupTest() {
	suite.testable = NewSubscribersReportsFilter()
}

func (suite *SalesSubscribersReportsFilterSuite) TestIsInvalidReportType() {
	suite.testable.SubTypeDetailed().Version12().Daily()
	suite.testable.TypePreOrder()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscribersReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func (suite *SalesSubscribersReportsFilterSuite) TestIsInvalidReportSubType() {
	suite.testable.Version12().Daily()
	suite.testable.SubTypeSummary()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscribersReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func (suite *SalesSubscribersReportsFilterSuite) TestIsInvalidReportFrequency() {
	suite.testable.SubTypeDetailed().Version12()
	suite.testable.Yearly()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscribersReportsFilter.IsValid: Frequency is not valid", err.Error())
}

func (suite *SalesSubscribersReportsFilterSuite) TestIsInvalidReportVersion() {
	suite.testable.SubTypeDetailed().Daily()
	suite.testable.Version10()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscribersReportsFilter.IsValid: Version is not valid", err.Error())
}

func TestSalesSubscribersReportsFilterSuite(t *testing.T) {
	suite.Run(t, new(SalesSubscribersReportsFilterSuite))
}

type SalesSubscriptionsOffersCodesRedemptionReportsFilterSuite struct {
	suite.Suite
	testable *SubscriptionsOffersCodesRedemptionReportsFilter
}

func (suite *SalesSubscriptionsOffersCodesRedemptionReportsFilterSuite) SetupTest() {
	suite.testable = NewSubscriptionsOffersCodesRedemptionReportsFilter()
}

func (suite *SalesSubscriptionsOffersCodesRedemptionReportsFilterSuite) TestIsInvalidReportType() {
	suite.testable.SubTypeDetailed().Version12().Daily()
	suite.testable.TypePreOrder()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func (suite *SalesSubscriptionsOffersCodesRedemptionReportsFilterSuite) TestIsInvalidReportSubType() {
	suite.testable.Version10().Daily()
	suite.testable.SubTypeDetailed()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func (suite *SalesSubscriptionsOffersCodesRedemptionReportsFilterSuite) TestIsInvalidReportFrequency() {
	suite.testable.SubTypeSummary().Version10()
	suite.testable.Yearly()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: Frequency is not valid", err.Error())
}

func (suite *SalesSubscriptionsOffersCodesRedemptionReportsFilterSuite) TestIsInvalidReportVersion() {
	suite.testable.SubTypeSummary().Daily()
	suite.testable.Version12()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: Version is not valid", err.Error())
}

func TestSalesSubscriptionsOffersCodesRedemptionReportsFilterSuite(t *testing.T) {
	suite.Run(t, new(SalesSubscriptionsOffersCodesRedemptionReportsFilterSuite))
}

type SalesNewsstandReportsFilterSuite struct {
	suite.Suite
	testable *NewsstandReportsFilter
}

func (suite *SalesNewsstandReportsFilterSuite) SetupTest() {
	suite.testable = NewNewsstandReportsFilter()
}

func (suite *SalesNewsstandReportsFilterSuite) TestIsInvalidReportType() {
	suite.testable.SubTypeDetailed().Version10().Daily()
	suite.testable.TypePreOrder()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "NewsstandReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func (suite *SalesNewsstandReportsFilterSuite) TestIsInvalidReportSubType() {
	suite.testable.Version10().Daily()
	suite.testable.SubTypeSummary()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "NewsstandReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func (suite *SalesNewsstandReportsFilterSuite) TestIsInvalidReportFrequency() {
	suite.testable.SubTypeDetailed().Version10()
	suite.testable.Yearly()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "NewsstandReportsFilter.IsValid: Frequency is not valid", err.Error())
}

func (suite *SalesNewsstandReportsFilterSuite) TestIsInvalidReportVersion() {
	suite.testable.SubTypeDetailed().Daily()
	suite.testable.Version12()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "NewsstandReportsFilter.IsValid: Version is not valid", err.Error())
}

func TestSalesNewsstandReportsFilterSuite(t *testing.T) {
	suite.Run(t, new(SalesNewsstandReportsFilterSuite))
}

type SalesPreOrdersReportsFilterSuite struct {
	suite.Suite
	testable *PreOrdersReportsFilter
}

func (suite *SalesPreOrdersReportsFilterSuite) SetupTest() {
	suite.testable = NewPreOrdersReportsFilter()
}

func (suite *SalesPreOrdersReportsFilterSuite) TestIsInvalidReportType() {
	suite.testable.SubTypeSummary().Version10().Daily()
	suite.testable.TypeNewsStand()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "PreOrdersReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func (suite *SalesPreOrdersReportsFilterSuite) TestIsInvalidReportSubType() {
	suite.testable.Version10().Daily()
	suite.testable.SubTypeDetailed()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "PreOrdersReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func (suite *SalesPreOrdersReportsFilterSuite) TestIsInvalidReportVersion() {
	suite.testable.SubTypeSummary().Daily()
	suite.testable.Version12()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "PreOrdersReportsFilter.IsValid: Version is not valid", err.Error())
}

func TestSalesPreOrdersReportsFilterSuite(t *testing.T) {
	suite.Run(t, new(SalesPreOrdersReportsFilterSuite))
}
