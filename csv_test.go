package appstore

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func Test_CSV_NewCSVReader(t *testing.T) {
	result := NewCSVReader(bytes.NewReader([]byte("")))
	assert.NotEmpty(t, result)
}

func Test_CSV_UnmarshalCSVSalesReportSales(t *testing.T) {
	reportData, _ := ioutil.ReadFile("stubs/reports/sales/sales.tsv")
	reports := []*SalesReportSale{}
	_ = UnmarshalCSV(reportData, &reports)
	assert.Equal(t, 1234567890, reports[0].AppleIdentifier.Value())
	assert.Equal(t, "2020-10-05", reports[0].BeginDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, "2020-10-05", reports[0].EndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, float64(299), reports[0].CustomerPrice.Value())
	assert.Equal(t, 209.3000030517578, reports[0].DeveloperProceeds.Value())
	assert.Equal(t, 12, reports[0].Units.Value())
}

func Test_CSV_UnmarshalSalesReportSubscription(t *testing.T) {
	reportData, _ := ioutil.ReadFile("stubs/reports/sales/subscription.tsv")
	reports := []*SalesReportSubscription{}
	_ = UnmarshalCSV(reportData, &reports)
	assert.Equal(t, "FooBarApp", reports[0].AppName)
	assert.Equal(t, 1234567890, reports[0].AppAppleID.Value())
	assert.Equal(t, 1234567890, reports[0].SubscriptionAppleID.Value())
	assert.Equal(t, 1234567890, reports[0].SubscriptionGroupID.Value())
	assert.Equal(t, float64(2950), reports[0].CustomerPrice.Value())
	assert.Equal(t, float64(2065), reports[0].DeveloperProceeds.Value())
	assert.Equal(t, 0, reports[0].BillingRetry.Value())
	assert.Equal(t, 20, reports[0].ActiveStandardPriceSubscriptions.Value())
	assert.Equal(t, 0, reports[0].ActiveFreeTrialIntroductoryOfferSubscriptions.Value())
	assert.Equal(t, 0, reports[0].ActivePayUpFrontIntroductoryOfferSubscriptions.Value())
	assert.Equal(t, 0, reports[0].ActivePayAsYouGoIntroductoryOfferSubscriptions.Value())
	assert.Equal(t, 0, reports[0].FreeTrialPromotionalOfferSubscriptions.Value())
	assert.Equal(t, 0, reports[0].PayUpFrontPromotionalOfferSubscriptions.Value())
	assert.Equal(t, 0, reports[0].PayAsYouGoPromotionalOfferSubscriptions.Value())
	assert.Equal(t, 0, reports[0].MarketingOptIns.Value())
	assert.Equal(t, 0, reports[0].GracePeriod.Value())
}

func Test_CSV_UnmarshalSalesReportSubscriptionEvent(t *testing.T) {
	reportData, _ := ioutil.ReadFile("stubs/reports/sales/subscription-event.tsv")
	reports := []*SalesReportSubscriptionEvent{}
	_ = UnmarshalCSV(reportData, &reports)
	assert.Equal(t, 1234567890, reports[0].AppAppleID.Value())
	assert.Equal(t, 1234567890, reports[0].SubscriptionAppleID.Value())
	assert.Equal(t, 1234567890, reports[0].SubscriptionGroupID.Value())
	assert.Equal(t, "2020-07-25", reports[0].OriginalStartDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, 11, reports[0].ConsecutivePaidPeriods.Value())
	assert.Equal(t, 1, reports[0].Quantity.Value())
	assert.Equal(t, "2020-10-06", reports[0].EventDate.Value().Format(CustomDateFormatDefault))
}

func Test_CSV_UnmarshalSalesReportSubscriber(t *testing.T) {
	reportData, _ := ioutil.ReadFile("stubs/reports/sales/subscriber.tsv")
	reports := []*SalesReportSubscriber{}
	_ = UnmarshalCSV(reportData, &reports)
	assert.Equal(t, 1234567890, reports[0].AppAppleID.Value())
	assert.Equal(t, 1234567890, reports[0].SubscriptionAppleID.Value())
	assert.Equal(t, 1234567890, reports[0].SubscriptionGroupID.Value())
	assert.Equal(t, "2020-10-05", reports[0].EventDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, 4.489999771118164, reports[0].CustomerPrice.Value())
	assert.Equal(t, 3.1500000953674316, reports[0].DeveloperProceeds.Value())
	assert.Equal(t, 1, reports[0].Units.Value())
	assert.True(t, reports[0].PurchaseDate.Value().IsZero())
}
