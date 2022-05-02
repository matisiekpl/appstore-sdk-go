package appstore

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Sales_SalesReportsBaseFilter_FillByDefault(t *testing.T) {
	filter := &SalesReportsBaseFilter{}
	filter.Daily().TypeSales().SubTypeSummary().Version10()
	assert.Equal(t, SalesReportTypeSales, filter.ReportType)
	assert.Equal(t, SalesReportSubTypeSummary, filter.ReportSubType)
	assert.Equal(t, SalesReportFrequencyDaily, filter.Frequency)
	assert.Equal(t, SalesReportVersion10, filter.Version)
}

func Test_Sales_SalesReportsBaseFilter_ToQueryParamsMapOnlyRequired(t *testing.T) {
	filter := &SalesReportsBaseFilter{}
	filter.Yearly().TypeSales().SubTypeSummary()

	qs := make(map[string]interface{})
	qs["filter[reportSubType]"] = string(SalesReportSubTypeSummary)
	qs["filter[reportType]"] = string(SalesReportTypeSales)
	qs["filter[frequency]"] = string(SalesReportFrequencyYearly)
	assert.Equal(t, qs, filter.ToQueryParamsMap())
}

func Test_Sales_SalesReportsBaseFilter_ToQueryParamsMap(t *testing.T) {
	filter := &SalesReportsBaseFilter{}
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	filter.Daily().TypeSales().SubTypeSummary().Version10().SetReportDate(date)

	qs := make(map[string]interface{})
	qs["filter[reportDate]"] = "2020-05-05"
	qs["filter[reportSubType]"] = string(SalesReportSubTypeSummary)
	qs["filter[reportType]"] = string(SalesReportTypeSales)
	qs["filter[frequency]"] = string(SalesReportFrequencyDaily)
	qs["filter[version]"] = string(SalesReportVersion10)
	assert.Equal(t, qs, filter.ToQueryParamsMap())
}

func Test_Sales_SalesReportsBaseFilter_SetSubType(t *testing.T) {
	filter := &SalesReportsBaseFilter{}
	filter.SubTypeSummary()
	assert.Equal(t, filter.ReportSubType, SalesReportSubTypeSummary)
	filter.SubTypeDetailed()
	assert.Equal(t, filter.ReportSubType, SalesReportSubTypeDetailed)
	filter.SubTypeOptIn()
	assert.Equal(t, filter.ReportSubType, SalesReportSubTypeOptIn)
}

func Test_Sales_SalesReportsBaseFilter_SetType(t *testing.T) {
	filter := &SalesReportsBaseFilter{}
	filter.TypeSales()
	assert.Equal(t, filter.ReportType, SalesReportTypeSales)
	filter.TypeNewsStand()
	assert.Equal(t, filter.ReportType, SalesReportTypeNewsStand)
	filter.TypePreOrder()
	assert.Equal(t, filter.ReportType, SalesReportTypePreorder)
	filter.TypeSubscriber()
	assert.Equal(t, filter.ReportType, SalesReportTypeSubscriber)
	filter.TypeSubscription()
	assert.Equal(t, filter.ReportType, SalesReportTypeSubscription)
	filter.TypeSubscriptionEvent()
	assert.Equal(t, filter.ReportType, SalesReportTypeSubscriptionEvent)
}

func Test_Sales_SalesReportsBaseFilter_SetFrequency(t *testing.T) {
	filter := &SalesReportsBaseFilter{}
	filter.Daily()
	assert.Equal(t, filter.Frequency, SalesReportFrequencyDaily)
	filter.Weekly()
	assert.Equal(t, filter.Frequency, SalesReportFrequencyWeekly)
	filter.Monthly()
	assert.Equal(t, filter.Frequency, SalesReportFrequencyMonthly)
	filter.Yearly()
	assert.Equal(t, filter.Frequency, SalesReportFrequencyYearly)
}

func Test_Sales_SalesReportsBaseFilter_SetVersion(t *testing.T) {
	filter := &SalesReportsBaseFilter{}
	filter.Version10()
	assert.Equal(t, filter.Version, SalesReportVersion10)
	filter.Version12()
	assert.Equal(t, filter.Version, SalesReportVersion12)
	filter.Version13()
	assert.Equal(t, filter.Version, SalesReportVersion13)
}

func Test_Sales_SalesReportsBaseFilter_IsValid(t *testing.T) {
	filter := &SalesReportsBaseFilter{}
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	filter.Daily().TypeSales().SubTypeSummary().Version10().SetReportDate(date)
	err := filter.IsValid()
	assert.Nil(t, err)
}

func Test_Sales_SalesReportsBaseFilter_IsInvalidReportType(t *testing.T) {
	filter := &SalesReportsBaseFilter{}
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	filter.Daily().SubTypeSummary().Version10().SetReportDate(date)
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SalesReportsBaseFilter.IsValid: ReportType is required", err.Error())
}

func Test_Sales_SalesReportsBaseFilter_IsInvalidReportSubType(t *testing.T) {
	filter := &SalesReportsBaseFilter{}
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	filter.Daily().TypeSales().Version10().SetReportDate(date)
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SalesReportsBaseFilter.IsValid: ReportSubType is required", err.Error())
}

func Test_Sales_SalesReportsBaseFilter_IsInvalidFrequency(t *testing.T) {
	filter := &SalesReportsBaseFilter{}
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	filter.SubTypeSummary().TypeSales().Version10().SetReportDate(date)
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SalesReportsBaseFilter.IsValid: Frequency is required", err.Error())
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
