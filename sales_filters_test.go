package appstore

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type SalesReportsBaseFilterTestSuite struct {
	suite.Suite
	testable *SalesReportsBaseFilter
}

func (suite *SalesReportsBaseFilterTestSuite) SetupTest() {
	suite.testable = &SalesReportsBaseFilter{}
}

func (suite *SalesReportsBaseFilterTestSuite) TestFillByDefault() {
	suite.testable.Daily().TypeSales().SubTypeSummary().Version10()
	assert.Equal(suite.T(), SalesReportTypeSales, suite.testable.ReportType)
	assert.Equal(suite.T(), SalesReportSubTypeSummary, suite.testable.ReportSubType)
	assert.Equal(suite.T(), SalesReportFrequencyDaily, suite.testable.Frequency)
	assert.Equal(suite.T(), SalesReportVersion10, suite.testable.Version)
}

func (suite *SalesReportsBaseFilterTestSuite) TestToQueryParamsMapOnlyRequired() {
	suite.testable.Yearly().TypeSales().SubTypeSummary()
	qs := make(map[string]interface{})
	qs["filter[reportSubType]"] = string(SalesReportSubTypeSummary)
	qs["filter[reportType]"] = string(SalesReportTypeSales)
	qs["filter[frequency]"] = string(SalesReportFrequencyYearly)
	assert.Equal(suite.T(), qs, suite.testable.ToQueryParamsMap())
}

func (suite *SalesReportsBaseFilterTestSuite) TestToQueryParamsMap() {
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

func (suite *SalesReportsBaseFilterTestSuite) TestSetSubType() {
	suite.testable.SubTypeSummary()
	assert.Equal(suite.T(), suite.testable.ReportSubType, SalesReportSubTypeSummary)
	suite.testable.SubTypeDetailed()
	assert.Equal(suite.T(), suite.testable.ReportSubType, SalesReportSubTypeDetailed)
	suite.testable.SubTypeOptIn()
	assert.Equal(suite.T(), suite.testable.ReportSubType, SalesReportSubTypeOptIn)
}

func (suite *SalesReportsBaseFilterTestSuite) TestSetType() {
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

func (suite *SalesReportsBaseFilterTestSuite) TestSetFrequency() {
	suite.testable.Daily()
	assert.Equal(suite.T(), suite.testable.Frequency, SalesReportFrequencyDaily)
	suite.testable.Weekly()
	assert.Equal(suite.T(), suite.testable.Frequency, SalesReportFrequencyWeekly)
	suite.testable.Monthly()
	assert.Equal(suite.T(), suite.testable.Frequency, SalesReportFrequencyMonthly)
	suite.testable.Yearly()
	assert.Equal(suite.T(), suite.testable.Frequency, SalesReportFrequencyYearly)
}

func (suite *SalesReportsBaseFilterTestSuite) TestSetVersion() {
	suite.testable.Version10()
	assert.Equal(suite.T(), suite.testable.Version, SalesReportVersion10)
	suite.testable.Version12()
	assert.Equal(suite.T(), suite.testable.Version, SalesReportVersion12)
	suite.testable.Version13()
	assert.Equal(suite.T(), suite.testable.Version, SalesReportVersion13)
}

func (suite *SalesReportsBaseFilterTestSuite) TestIsValid() {
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	suite.testable.Daily().TypeSales().SubTypeSummary().Version10().SetReportDate(date)
	err := suite.testable.IsValid()
	assert.Nil(suite.T(), err)
}

func (suite *SalesReportsBaseFilterTestSuite) TestIsInvalidReportType() {
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	suite.testable.Daily().SubTypeSummary().Version10().SetReportDate(date)
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SalesReportsBaseFilter.IsValid: ReportType is required", err.Error())
}

func (suite *SalesReportsBaseFilterTestSuite) TestIsInvalidReportSubType() {
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	suite.testable.Daily().TypeSales().Version10().SetReportDate(date)
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SalesReportsBaseFilter.IsValid: ReportSubType is required", err.Error())
}

func (suite *SalesReportsBaseFilterTestSuite) TestIsInvalidFrequency() {
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	suite.testable.SubTypeSummary().TypeSales().Version10().SetReportDate(date)
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SalesReportsBaseFilter.IsValid: Frequency is required", err.Error())
}

func TestSalesReportsBaseFilterTestSuite(t *testing.T) {
	suite.Run(t, new(SalesReportsBaseFilterTestSuite))
}

type SalesReportsFilterTestSuite struct {
	suite.Suite
	testable *SalesReportsFilter
}

func (suite *SalesReportsFilterTestSuite) SetupTest() {
	suite.testable = NewSalesReportsFilter()
}

func (suite *SalesReportsFilterTestSuite) TestIsInvalidReportType() {
	suite.testable.SubTypeSummary().Version10().Daily()
	suite.testable.TypePreOrder()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SalesReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func (suite *SalesReportsFilterTestSuite) TestIsInvalidReportSubType() {
	suite.testable.Version10().Daily()
	suite.testable.SubTypeDetailed()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SalesReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func (suite *SalesReportsFilterTestSuite) TestIsInvalidVersion() {
	suite.testable.SubTypeSummary().Daily()
	suite.testable.Version12()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SalesReportsFilter.IsValid: Version is not valid", err.Error())
}

func TestSalesReportsFilterTestSuite(t *testing.T) {
	suite.Run(t, new(SalesReportsFilterTestSuite))
}

type SalesSubscriptionsReportsFilterTestSuite struct {
	suite.Suite
	testable *SubscriptionsReportsFilter
}

func (suite *SalesSubscriptionsReportsFilterTestSuite) SetupTest() {
	suite.testable = NewSubscriptionsReportsFilter()
}

func (suite *SalesSubscriptionsReportsFilterTestSuite) TestIsInvalidReportType() {
	suite.testable.SubTypeSummary().Version12().Daily()
	suite.testable.TypePreOrder()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func (suite *SalesSubscriptionsReportsFilterTestSuite) TestIsInvalidReportSubType() {
	suite.testable.Version12().Daily()
	suite.testable.SubTypeOptIn()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func (suite *SalesSubscriptionsReportsFilterTestSuite) TestIsInvalidFrequency() {
	suite.testable.SubTypeSummary().Version12()
	suite.testable.Yearly()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsReportsFilter.IsValid: Frequency is not valid", err.Error())
}

func (suite *SalesSubscriptionsReportsFilterTestSuite) TestIsInvalidVersion() {
	suite.testable.SubTypeSummary().Daily()
	suite.testable.Version10()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsReportsFilter.IsValid: Version is not valid", err.Error())
}

func TestSalesSubscriptionsReportsFilterTestSuite(t *testing.T) {
	suite.Run(t, new(SalesSubscriptionsReportsFilterTestSuite))
}

type SalesSubscriptionsEventsReportsFilterTestSuite struct {
	suite.Suite
	testable *SubscriptionsEventsReportsFilter
}

func (suite *SalesSubscriptionsEventsReportsFilterTestSuite) SetupTest() {
	suite.testable = NewSubscriptionsEventsReportsFilter()
}

func (suite *SalesSubscriptionsEventsReportsFilterTestSuite) TestIsInvalidReportType() {
	suite.testable.SubTypeSummary().Version12().Daily()
	suite.testable.TypePreOrder()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsEventsReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func (suite *SalesSubscriptionsEventsReportsFilterTestSuite) TestIsInvalidReportSubType() {
	suite.testable.Version12().Daily()
	suite.testable.SubTypeOptIn()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsEventsReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func (suite *SalesSubscriptionsEventsReportsFilterTestSuite) TestIsInvalidReportFrequency() {
	suite.testable.SubTypeSummary().Version12()
	suite.testable.Yearly()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsEventsReportsFilter.IsValid: Frequency is not valid", err.Error())
}

func (suite *SalesSubscriptionsEventsReportsFilterTestSuite) TestIsInvalidReportVersion() {
	suite.testable.SubTypeSummary().Daily()
	suite.testable.Version10()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsEventsReportsFilter.IsValid: Version is not valid", err.Error())
}

func TestSalesSubscriptionsEventsReportsFilterTestSuite(t *testing.T) {
	suite.Run(t, new(SalesSubscriptionsEventsReportsFilterTestSuite))
}

type SalesSubscribersReportsFilterTestSuite struct {
	suite.Suite
	testable *SubscribersReportsFilter
}

func (suite *SalesSubscribersReportsFilterTestSuite) SetupTest() {
	suite.testable = NewSubscribersReportsFilter()
}

func (suite *SalesSubscribersReportsFilterTestSuite) TestIsInvalidReportType() {
	suite.testable.SubTypeDetailed().Version12().Daily()
	suite.testable.TypePreOrder()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscribersReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func (suite *SalesSubscribersReportsFilterTestSuite) TestIsInvalidReportSubType() {
	suite.testable.Version12().Daily()
	suite.testable.SubTypeSummary()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscribersReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func (suite *SalesSubscribersReportsFilterTestSuite) TestIsInvalidReportFrequency() {
	suite.testable.SubTypeDetailed().Version12()
	suite.testable.Yearly()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscribersReportsFilter.IsValid: Frequency is not valid", err.Error())
}

func (suite *SalesSubscribersReportsFilterTestSuite) TestIsInvalidReportVersion() {
	suite.testable.SubTypeDetailed().Daily()
	suite.testable.Version10()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscribersReportsFilter.IsValid: Version is not valid", err.Error())
}

func TestSalesSubscribersReportsFilterTestSuite(t *testing.T) {
	suite.Run(t, new(SalesSubscribersReportsFilterTestSuite))
}

type SalesSubscriptionsOffersCodesRedemptionReportsFilterTestSuite struct {
	suite.Suite
	testable *SubscriptionsOffersCodesRedemptionReportsFilter
}

func (suite *SalesSubscriptionsOffersCodesRedemptionReportsFilterTestSuite) SetupTest() {
	suite.testable = NewSubscriptionsOffersCodesRedemptionReportsFilter()
}

func (suite *SalesSubscriptionsOffersCodesRedemptionReportsFilterTestSuite) TestIsInvalidReportType() {
	suite.testable.SubTypeDetailed().Version12().Daily()
	suite.testable.TypePreOrder()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func (suite *SalesSubscriptionsOffersCodesRedemptionReportsFilterTestSuite) TestIsInvalidReportSubType() {
	suite.testable.Version10().Daily()
	suite.testable.SubTypeDetailed()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func (suite *SalesSubscriptionsOffersCodesRedemptionReportsFilterTestSuite) TestIsInvalidReportFrequency() {
	suite.testable.SubTypeSummary().Version10()
	suite.testable.Yearly()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: Frequency is not valid", err.Error())
}

func (suite *SalesSubscriptionsOffersCodesRedemptionReportsFilterTestSuite) TestIsInvalidReportVersion() {
	suite.testable.SubTypeSummary().Daily()
	suite.testable.Version12()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: Version is not valid", err.Error())
}

func TestSalesSubscriptionsOffersCodesRedemptionReportsFilterTestSuite(t *testing.T) {
	suite.Run(t, new(SalesSubscriptionsOffersCodesRedemptionReportsFilterTestSuite))
}

type SalesNewsstandReportsFilterTestSuite struct {
	suite.Suite
	testable *NewsstandReportsFilter
}

func (suite *SalesNewsstandReportsFilterTestSuite) SetupTest() {
	suite.testable = NewNewsstandReportsFilter()
}

func (suite *SalesNewsstandReportsFilterTestSuite) TestIsInvalidReportType() {
	suite.testable.SubTypeDetailed().Version10().Daily()
	suite.testable.TypePreOrder()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "NewsstandReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func (suite *SalesNewsstandReportsFilterTestSuite) TestIsInvalidReportSubType() {
	suite.testable.Version10().Daily()
	suite.testable.SubTypeSummary()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "NewsstandReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func (suite *SalesNewsstandReportsFilterTestSuite) TestIsInvalidReportFrequency() {
	suite.testable.SubTypeDetailed().Version10()
	suite.testable.Yearly()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "NewsstandReportsFilter.IsValid: Frequency is not valid", err.Error())
}

func (suite *SalesNewsstandReportsFilterTestSuite) TestIsInvalidReportVersion() {
	suite.testable.SubTypeDetailed().Daily()
	suite.testable.Version12()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "NewsstandReportsFilter.IsValid: Version is not valid", err.Error())
}

func TestSalesNewsstandReportsFilterTestSuite(t *testing.T) {
	suite.Run(t, new(SalesNewsstandReportsFilterTestSuite))
}

type SalesPreOrdersReportsFilterTestSuite struct {
	suite.Suite
	testable *PreOrdersReportsFilter
}

func (suite *SalesPreOrdersReportsFilterTestSuite) SetupTest() {
	suite.testable = NewPreOrdersReportsFilter()
}

func (suite *SalesPreOrdersReportsFilterTestSuite) TestIsInvalidReportType() {
	suite.testable.SubTypeSummary().Version10().Daily()
	suite.testable.TypeNewsStand()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "PreOrdersReportsFilter.IsValid: ReportType is not valid", err.Error())
}

func (suite *SalesPreOrdersReportsFilterTestSuite) TestIsInvalidReportSubType() {
	suite.testable.Version10().Daily()
	suite.testable.SubTypeDetailed()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "PreOrdersReportsFilter.IsValid: ReportSubType is not valid", err.Error())
}

func (suite *SalesPreOrdersReportsFilterTestSuite) TestIsInvalidReportVersion() {
	suite.testable.SubTypeSummary().Daily()
	suite.testable.Version12()
	err := suite.testable.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "PreOrdersReportsFilter.IsValid: Version is not valid", err.Error())
}

func TestSalesPreOrdersReportsFilterTestSuite(t *testing.T) {
	suite.Run(t, new(SalesPreOrdersReportsFilterTestSuite))
}
