package appstore

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Finances_FinancesReportsFilter_IsValid(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2020-04-17")
	filter := NewFinancesReportsFilter()
	filter.SetReportDate(date).SetRegionCode("US")
	assert.NoError(t, filter.IsValid())
}

func Test_Finances_FinancesReportsFilter_IsInvalidEmptyReportType(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2020-04-17")
	filter := &FinancesReportsFilter{ReportDate: date, RegionCode: "US"}
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "FinancesReportsFilter.IsValid: ReportType is required", err.Error())
}

func Test_Finances_FinancesReportsFilter_IsInvalidEmptyRegionCode(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2020-04-17")
	filter := &FinancesReportsFilter{ReportDate: date}
	filter.TypeFinancial()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "FinancesReportsFilter.IsValid: RegionCode is required", err.Error())
}

func Test_Finances_FinancesReportsFilter_IsInvalidEmptyReportDate(t *testing.T) {
	filter := &FinancesReportsFilter{RegionCode: "US"}
	filter.TypeFinanceDetail()
	err := filter.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "FinancesReportsFilter.IsValid: ReportDate is required", err.Error())
}

func Test_Finances_FinancesReportsFilter_ToQueryParamsMap(t *testing.T) {
	date, _ := time.Parse("2006-01-02", "2020-05-04")
	filter := &FinancesReportsFilter{ReportDate: date, RegionCode: "US", ReportType: FinancesReportTypeFinancial}
	qs := make(map[string]interface{})
	qs["filter[reportDate]"] = "2020-05"
	qs["filter[reportType]"] = string(FinancesReportTypeFinancial)
	qs["filter[regionCode]"] = "US"
	assert.Equal(t, qs, filter.toQueryParamsMap())
}
