package appstore_sdk

import (
	"fmt"
	"time"
)

type SalesReportType string
type SalesReportSubType string
type SalesReportFrequency string
type SalesReportVersion string

const (
	SalesReportFrequencyDaily   SalesReportFrequency = "DAILY"
	SalesReportFrequencyWeekly  SalesReportFrequency = "WEEKLY"
	SalesReportFrequencyMonthly SalesReportFrequency = "MONTHLY"
	SalesReportFrequencyYearly  SalesReportFrequency = "YEARLY"
)

const (
	SalesReportTypeSales             SalesReportType = "SALES"
	SalesReportTypePreorder          SalesReportType = "PRE_ORDER"
	SalesReportTypeNewsStand         SalesReportType = "NEWSSTAND"
	SalesReportTypeSubscription      SalesReportType = "SUBSCRIPTION"
	SalesReportTypeSubscriptionEvent SalesReportType = "SUBSCRIPTION_EVENT"
	SalesReportTypeSubscriber        SalesReportType = "SUBSCRIBER"
)

const (
	SalesReportSubTypeSummary  SalesReportSubType = "SUMMARY"
	SalesReportSubTypeDetailed SalesReportSubType = "DETAILED"
	SalesReportSubTypeOptIn    SalesReportSubType = "OPT_IN"
)

const (
	SalesReportVersion10 SalesReportVersion = "1_0"
	SalesReportVersion12 SalesReportVersion = "1_2"
)

type SalesReportsFilter struct {
	ReportDate    time.Time
	ReportSubType SalesReportSubType
	ReportType    SalesReportType
	Frequency     SalesReportFrequency
	Version       SalesReportVersion
	VendorNumber  string
}

func (f *SalesReportsFilter) SetReportDate(value time.Time) *SalesReportsFilter {
	f.ReportDate = value
	return f
}

func (f *SalesReportsFilter) SetReportSubType(value SalesReportSubType) *SalesReportsFilter {
	f.ReportSubType = value
	return f
}

func (f *SalesReportsFilter) SubTypeDetailed() *SalesReportsFilter {
	return f.SetReportSubType(SalesReportSubTypeDetailed)
}

func (f *SalesReportsFilter) SubTypeOptIn() *SalesReportsFilter {
	return f.SetReportSubType(SalesReportSubTypeOptIn)
}

func (f *SalesReportsFilter) SubTypeSummary() *SalesReportsFilter {
	return f.SetReportSubType(SalesReportSubTypeSummary)
}

func (f *SalesReportsFilter) SetReportType(value SalesReportType) *SalesReportsFilter {
	f.ReportType = value
	return f
}

func (f *SalesReportsFilter) TypeSales() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypeSales)
}

func (f *SalesReportsFilter) TypePreOrder() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypePreorder)
}

func (f *SalesReportsFilter) TypeNewsStand() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypeNewsStand)
}

func (f *SalesReportsFilter) TypeSubscription() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypeSubscription)
}

func (f *SalesReportsFilter) TypeSubscriptionEvent() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypeSubscriptionEvent)
}

func (f *SalesReportsFilter) TypeSubscriber() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypeSubscriber)
}

func (f *SalesReportsFilter) SetFrequency(value SalesReportFrequency) *SalesReportsFilter {
	f.Frequency = value
	return f
}

func (f *SalesReportsFilter) Daily() *SalesReportsFilter {
	return f.SetFrequency(SalesReportFrequencyDaily)
}

func (f *SalesReportsFilter) Weekly() *SalesReportsFilter {
	return f.SetFrequency(SalesReportFrequencyWeekly)
}

func (f *SalesReportsFilter) Monthly() *SalesReportsFilter {
	return f.SetFrequency(SalesReportFrequencyMonthly)
}

func (f *SalesReportsFilter) Yearly() *SalesReportsFilter {
	return f.SetFrequency(SalesReportFrequencyYearly)
}

func (f *SalesReportsFilter) SetVersion(value SalesReportVersion) *SalesReportsFilter {
	f.Version = value
	return f
}

func (f *SalesReportsFilter) Version12() *SalesReportsFilter {
	return f.SetVersion(SalesReportVersion12)
}

func (f *SalesReportsFilter) Version10() *SalesReportsFilter {
	return f.SetVersion(SalesReportVersion10)
}

func (f *SalesReportsFilter) ToQueryParamsMap() map[string]interface{} {
	qs := make(map[string]interface{})
	qs["filter[reportDate]"] = f.ReportDate.Format("2006-01-02")
	qs["filter[reportSubType]"] = string(f.ReportSubType)
	qs["filter[reportType]"] = string(f.ReportType)
	qs["filter[frequency]"] = string(f.Frequency)
	qs["filter[version]"] = string(f.Version)
	qs["filter[vendorNumber]"] = f.VendorNumber
	return qs
}

func (f *SalesReportsFilter) IsValid() error {
	if f.ReportDate.IsZero() {
		return fmt.Errorf("SalesReportsFilter@IsValid: %v", "ReportDate is required")
	}
	if f.ReportType == "" {
		return fmt.Errorf("SalesReportsFilter@IsValid: %v", "ReportType is required")
	}
	if f.ReportSubType == "" {
		return fmt.Errorf("SalesReportsFilter@IsValid: %v", "ReportSubType is required")
	}
	if f.Frequency == "" {
		return fmt.Errorf("SalesReportsFilter@IsValid: %v", "Frequency is required")
	}
	if f.VendorNumber == "" {
		return fmt.Errorf("SalesReportsFilter@IsValid: %v", "VendorNumber is required")
	}
	return nil
}
