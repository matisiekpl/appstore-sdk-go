package appstore

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"testing"
)

type CSVTestSuite struct {
	suite.Suite
}

func (suite *CSVTestSuite) TestNewCSVReader() {
	result := NewCSVReader(bytes.NewReader([]byte("")))
	assert.NotEmpty(suite.T(), result)
}

func (suite *CSVTestSuite) TestUnmarshalCSVSalesReport() {
	reportData, _ := ioutil.ReadFile("stubs/reports/sales/sales.tsv")
	reports := []*SalesReport{}
	err := UnmarshalCSV(reportData, &reports)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1234567890, reports[0].AppleIdentifier.Value())
	assert.Equal(suite.T(), "2020-10-05", reports[0].BeginDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), "2020-10-05", reports[0].EndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), float64(299), reports[0].CustomerPrice.Value())
	assert.Equal(suite.T(), 209.3000030517578, reports[0].DeveloperProceeds.Value())
	assert.Equal(suite.T(), float64(12), reports[0].Units.Value())
}

func (suite *CSVTestSuite) TestUnmarshalSalesSubscriptionsReport() {
	reportData, _ := ioutil.ReadFile("stubs/reports/sales/subscriptions.tsv")
	reports := []*SubscriptionsReport{}
	err := UnmarshalCSV(reportData, &reports)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "FooBarApp", reports[0].AppName)
	assert.Equal(suite.T(), 1234567890, reports[0].AppAppleID.Value())
	assert.Equal(suite.T(), 1234567890, reports[0].SubscriptionAppleID.Value())
	assert.Equal(suite.T(), 1234567890, reports[0].SubscriptionGroupID.Value())
	assert.Equal(suite.T(), float64(2950), reports[0].CustomerPrice.Value())
	assert.Equal(suite.T(), float64(2065), reports[0].DeveloperProceeds.Value())
	assert.Equal(suite.T(), 0, reports[0].BillingRetry.Value())
	assert.Equal(suite.T(), 20, reports[0].ActiveStandardPriceSubscriptions.Value())
	assert.Equal(suite.T(), 0, reports[0].ActiveFreeTrialIntroductoryOfferSubscriptions.Value())
	assert.Equal(suite.T(), 0, reports[0].ActivePayUpFrontIntroductoryOfferSubscriptions.Value())
	assert.Equal(suite.T(), 0, reports[0].ActivePayAsYouGoIntroductoryOfferSubscriptions.Value())
	assert.Equal(suite.T(), 0, reports[0].FreeTrialPromotionalOfferSubscriptions.Value())
	assert.Equal(suite.T(), 0, reports[0].PayUpFrontPromotionalOfferSubscriptions.Value())
	assert.Equal(suite.T(), 0, reports[0].PayAsYouGoPromotionalOfferSubscriptions.Value())
	assert.Equal(suite.T(), 0, reports[0].MarketingOptIns.Value())
	assert.Equal(suite.T(), 0, reports[0].GracePeriod.Value())
}

func (suite *CSVTestSuite) TestUnmarshalSalesSubscriptionsEventsReport() {
	reportData, _ := ioutil.ReadFile("stubs/reports/sales/subscriptions-events.tsv")
	reports := []*SubscriptionsEventsReport{}
	err := UnmarshalCSV(reportData, &reports)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1234567890, reports[0].AppAppleID.Value())
	assert.Equal(suite.T(), 1234567890, reports[0].SubscriptionAppleID.Value())
	assert.Equal(suite.T(), 1234567890, reports[0].SubscriptionGroupID.Value())
	assert.Equal(suite.T(), "2020-07-25", reports[0].OriginalStartDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), 11, reports[0].ConsecutivePaidPeriods.Value())
	assert.Equal(suite.T(), 1, reports[0].Quantity.Value())
	assert.Equal(suite.T(), "2020-10-06", reports[0].EventDate.Value().Format(CustomDateFormatDefault))
}

func (suite *CSVTestSuite) TestUnmarshalSalesSubscribersReport() {
	reportData, _ := ioutil.ReadFile("stubs/reports/sales/subscribers.tsv")
	reports := []*SubscribersReport{}
	err := UnmarshalCSV(reportData, &reports)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1234567890, reports[0].AppAppleID.Value())
	assert.Equal(suite.T(), 1234567890, reports[0].SubscriptionAppleID.Value())
	assert.Equal(suite.T(), 1234567890, reports[0].SubscriptionGroupID.Value())
	assert.Equal(suite.T(), "2020-10-05", reports[0].EventDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), 4.489999771118164, reports[0].CustomerPrice.Value())
	assert.Equal(suite.T(), 3.1500000953674316, reports[0].DeveloperProceeds.Value())
	assert.Equal(suite.T(), 1, reports[0].Units.Value())
	assert.True(suite.T(), reports[0].PurchaseDate.Value().IsZero())
}

func TestCSVTestSuite(t *testing.T) {
	suite.Run(t, new(CSVTestSuite))
}
