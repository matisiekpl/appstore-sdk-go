package appstore_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSalesReportsFilter_FillByDefault(t *testing.T) {
	filter := SalesReportsFilter{}
	filter.Daily().Sales().Summary().Version10()
	assert.Equal(t, ReportTypeSales, filter.ReportType)
	assert.Equal(t, ReportSubTypeSummary, filter.ReportSubType)
	assert.Equal(t, FrequencyDaily, filter.Frequency)
	assert.Equal(t, Version10, filter.Version)
}

func TestSalesReportsFilter_ToQueryParamsMap(t *testing.T) {
	filter := SalesReportsFilter{}
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	filter.Daily().Sales().Summary().Version10().SetReportDate(date)

	qs := make(map[string]string)
	qs["filter[reportDate]"] = "2020-05-05"
	qs["filter[reportSubType]"] = string(ReportSubTypeSummary)
	qs["filter[reportType]"] = string(ReportTypeSales)
	qs["filter[frequency]"] = string(FrequencyDaily)
	qs["filter[version]"] = string(Version10)
	qs["filter[vendorNumber]"] = ""
	assert.Equal(t, qs, filter.ToQueryParamsMap())
}
