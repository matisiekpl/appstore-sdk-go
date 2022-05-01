package appstore

import (
	"fmt"
	"time"
)

//FinancesReportType type
type FinancesReportType string

const (
	//FinancesReportTypeFinancial const
	FinancesReportTypeFinancial FinancesReportType = "FINANCIAL"
	//FinancesReportTypeFinanceDetail const
	FinancesReportTypeFinanceDetail FinancesReportType = "FINANCE_DETAIL"
)

//FinancesReportsFilter sales reports filter
type FinancesReportsFilter struct {
	ReportDate time.Time          //(Required) The fiscal month of the report you wish to download based on the Apple Fiscal Calendar. The fiscal month is specified in the YYYY-MM format.
	ReportType FinancesReportType //(Required) This value is always FINANCIAL. Possible values: FINANCIAL, FINANCE_DETAIL
	RegionCode string             //(Required) You can download consolidated or separate financial reports per territory. For a complete list of possible values, see Financial Report Regions and Currencies.
}

//ToQueryParamsMap Convert filter to query params
func (f *FinancesReportsFilter) ToQueryParamsMap() map[string]interface{} {
	qs := make(map[string]interface{})
	qs["filter[reportType]"] = string(f.ReportType)
	qs["filter[regionCode]"] = f.RegionCode
	if !f.ReportDate.IsZero() {
		qs["filter[reportDate]"] = f.ReportDate.Format("2006-01")
	}
	return qs
}

//IsValid Validate sales report filter params
func (f *FinancesReportsFilter) IsValid() error {
	if f.ReportType == "" {
		return fmt.Errorf("FinancesReportsFilter.IsValid: %v", "ReportType is required")
	}
	if f.RegionCode == "" {
		return fmt.Errorf("FinancesReportsFilter.IsValid: %v", "RegionCode is required")
	}
	if f.ReportDate.IsZero() {
		return fmt.Errorf("FinancesReportsFilter.IsValid: %v", "ReportDate is required")
	}
	return nil
}
