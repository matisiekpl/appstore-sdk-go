package appstore

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_SalesReports_Filter_FillByDefault(t *testing.T) {
	filter := &SalesReportsFilter{}
	filter.Daily().TypeSales().SubTypeSummary().Version10()
	assert.Equal(t, SalesReportTypeSales, filter.ReportType)
	assert.Equal(t, SalesReportSubTypeSummary, filter.ReportSubType)
	assert.Equal(t, SalesReportFrequencyDaily, filter.Frequency)
	assert.Equal(t, SalesReportVersion10, filter.Version)
}

func Test_SalesReports_Filter_ToQueryParamsMapOnlyRequired(t *testing.T) {
	filter := &SalesReportsFilter{}
	filter.Yearly().TypeSales().SubTypeSummary()

	qs := make(map[string]interface{})
	qs["filter[reportSubType]"] = string(SalesReportSubTypeSummary)
	qs["filter[reportType]"] = string(SalesReportTypeSales)
	qs["filter[frequency]"] = string(SalesReportFrequencyYearly)
	qs["filter[vendorNumber]"] = ""
	assert.Equal(t, qs, filter.ToQueryParamsMap())
}

func Test_SalesReports_Filter_ToQueryParamsMap(t *testing.T) {
	filter := &SalesReportsFilter{}
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	filter.Daily().TypeSales().SubTypeSummary().Version10().SetReportDate(date)

	qs := make(map[string]interface{})
	qs["filter[reportDate]"] = "2020-05-05"
	qs["filter[reportSubType]"] = string(SalesReportSubTypeSummary)
	qs["filter[reportType]"] = string(SalesReportTypeSales)
	qs["filter[frequency]"] = string(SalesReportFrequencyDaily)
	qs["filter[version]"] = string(SalesReportVersion10)
	qs["filter[vendorNumber]"] = ""
	assert.Equal(t, qs, filter.ToQueryParamsMap())
}

func Test_SalesReports_Filter_SetSubType(t *testing.T) {
	filter := &SalesReportsFilter{}
	filter.SubTypeSummary()
	assert.Equal(t, filter.ReportSubType, SalesReportSubTypeSummary)
	filter.SubTypeDetailed()
	assert.Equal(t, filter.ReportSubType, SalesReportSubTypeDetailed)
	filter.SubTypeOptIn()
	assert.Equal(t, filter.ReportSubType, SalesReportSubTypeOptIn)
}

func Test_SalesReports_Filter_SetType(t *testing.T) {
	filter := &SalesReportsFilter{}
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

func Test_SalesReports_Filter_SetFrequency(t *testing.T) {
	filter := &SalesReportsFilter{}
	filter.Daily()
	assert.Equal(t, filter.Frequency, SalesReportFrequencyDaily)
	filter.Weekly()
	assert.Equal(t, filter.Frequency, SalesReportFrequencyWeekly)
	filter.Monthly()
	assert.Equal(t, filter.Frequency, SalesReportFrequencyMonthly)
	filter.Yearly()
	assert.Equal(t, filter.Frequency, SalesReportFrequencyYearly)
}

func Test_SalesReports_Filter_SetVersion(t *testing.T) {
	filter := &SalesReportsFilter{}
	filter.Version10()
	assert.Equal(t, filter.Version, SalesReportVersion10)
	filter.Version12()
	assert.Equal(t, filter.Version, SalesReportVersion12)
}

func Test_SalesReports_Filter_IsValid(t *testing.T) {
	filter := &SalesReportsFilter{}
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	filter.Daily().TypeSales().SubTypeSummary().Version10().SetReportDate(date)
	err := filter.IsValid()
	assert.Nil(t, err)
}

func Test_SalesReports_Filter_IsInValidReportType(t *testing.T) {
	filter := &SalesReportsFilter{}
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	filter.Daily().SubTypeSummary().Version10().SetReportDate(date)
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SalesReportsFilter.IsValid: ReportType is required", err.Error())
}

func Test_SalesReports_Filter_IsInValidReportSubType(t *testing.T) {
	filter := &SalesReportsFilter{}
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	filter.Daily().TypeSales().Version10().SetReportDate(date)
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SalesReportsFilter.IsValid: ReportSubType is required", err.Error())
}

func Test_SalesReports_Filter_IsInValidFrequency(t *testing.T) {
	filter := &SalesReportsFilter{}
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	filter.SubTypeSummary().TypeSales().Version10().SetReportDate(date)
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "SalesReportsFilter.IsValid: Frequency is required", err.Error())
}
