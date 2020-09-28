package appstore_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSalesReportsFilter_FillByDefault(t *testing.T) {
	filter := SalesReportsFilter{}
	filter.Daily().Sales().Summary().V1_0()
	assert.Equal(t, SALES, filter.ReportType)
	assert.Equal(t, SUMMARY, filter.ReportSubType)
	assert.Equal(t, DAILY, filter.Frequency)
	assert.Equal(t, V1_0, filter.Version)
}

func TestSalesReportsFilter_ToQueryParamsMap(t *testing.T) {
	filter := SalesReportsFilter{}
	date, _ := time.Parse("2006-01-02", "2020-05-05")
	filter.Daily().Sales().Summary().V1_0().SetReportDate(date)

	qs := make(map[string]string)
	qs["filter[reportDate]"] = "2020-05-05"
	qs["filter[reportSubType]"] = string(SUMMARY)
	qs["filter[reportType]"] = string(SALES)
	qs["filter[frequency]"] = string(DAILY)
	qs["filter[version]"] = string(V1_0)
	qs["filter[vendorNumber]"] = ""
	assert.Equal(t, qs, filter.ToQueryParamsMap())
}
