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

//Sales reports filter
type SalesReportsFilter struct {
	ReportDate    time.Time            //The report date to download. The date is specified in the YYYY-MM-DD format for all report frequencies except DAILY, which does not require a specified date. For more information, see report availability and storage.
	ReportSubType SalesReportSubType   //(Required) The report sub type to download. For a list of values, see Allowed Values Based on Sales Report Type below. Possible values: SUMMARY, DETAILED, OPT_IN
	ReportType    SalesReportType      //(Required) The report to download. For more details on each report type see About Reports. Possible values: SALES, PRE_ORDER, NEWSSTAND, SUBSCRIPTION, SUBSCRIPTION_EVENT, SUBSCRIBER
	Frequency     SalesReportFrequency //(Required) Frequency of the report to download. For a list of values, see Allowed Values Based on Sales Report Type below. Possible values: DAILY, WEEKLY, MONTHLY, YEARLY
	Version       SalesReportVersion   //The version of the report. For a list of values, see Allowed Values Based on Sales Report Type below.
	VendorNumber  string               //(Required) You can find your vendor number in Payments and Financial Reports.
}

//Set report date
func (f *SalesReportsFilter) SetReportDate(value time.Time) *SalesReportsFilter {
	f.ReportDate = value
	return f
}

//Set report sub type
func (f *SalesReportsFilter) SetReportSubType(value SalesReportSubType) *SalesReportsFilter {
	f.ReportSubType = value
	return f
}

//Change report sub type to Detailed
func (f *SalesReportsFilter) SubTypeDetailed() *SalesReportsFilter {
	return f.SetReportSubType(SalesReportSubTypeDetailed)
}

//Change report sub type to OptIn
func (f *SalesReportsFilter) SubTypeOptIn() *SalesReportsFilter {
	return f.SetReportSubType(SalesReportSubTypeOptIn)
}

//Change report sub type to Summary
func (f *SalesReportsFilter) SubTypeSummary() *SalesReportsFilter {
	return f.SetReportSubType(SalesReportSubTypeSummary)
}

//Set report type
func (f *SalesReportsFilter) SetReportType(value SalesReportType) *SalesReportsFilter {
	f.ReportType = value
	return f
}

//Change report type to Sales
func (f *SalesReportsFilter) TypeSales() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypeSales)
}

//Change report type to PreOrder
func (f *SalesReportsFilter) TypePreOrder() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypePreorder)
}

//Change report type to NewsStand
func (f *SalesReportsFilter) TypeNewsStand() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypeNewsStand)
}

//Change report type to Subscription
func (f *SalesReportsFilter) TypeSubscription() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypeSubscription)
}

//Change report type to SubscriptionEvent
func (f *SalesReportsFilter) TypeSubscriptionEvent() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypeSubscriptionEvent)
}

//Change report type to Subscriber
func (f *SalesReportsFilter) TypeSubscriber() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypeSubscriber)
}

//Set frequency
func (f *SalesReportsFilter) SetFrequency(value SalesReportFrequency) *SalesReportsFilter {
	f.Frequency = value
	return f
}

//Change frequency to Daily
func (f *SalesReportsFilter) Daily() *SalesReportsFilter {
	return f.SetFrequency(SalesReportFrequencyDaily)
}

//Change frequency to Weekly
func (f *SalesReportsFilter) Weekly() *SalesReportsFilter {
	return f.SetFrequency(SalesReportFrequencyWeekly)
}

//Change frequency to Monthly
func (f *SalesReportsFilter) Monthly() *SalesReportsFilter {
	return f.SetFrequency(SalesReportFrequencyMonthly)
}

//Change frequency to Yearly
func (f *SalesReportsFilter) Yearly() *SalesReportsFilter {
	return f.SetFrequency(SalesReportFrequencyYearly)
}

//Set version
func (f *SalesReportsFilter) SetVersion(value SalesReportVersion) *SalesReportsFilter {
	f.Version = value
	return f
}

//Change version to 1_2
func (f *SalesReportsFilter) Version12() *SalesReportsFilter {
	return f.SetVersion(SalesReportVersion12)
}

//Change version to 1_0
func (f *SalesReportsFilter) Version10() *SalesReportsFilter {
	return f.SetVersion(SalesReportVersion10)
}

//Convert filter to query params
func (f *SalesReportsFilter) ToQueryParamsMap() map[string]interface{} {
	qs := make(map[string]interface{})
	qs["filter[reportSubType]"] = string(f.ReportSubType)
	qs["filter[reportType]"] = string(f.ReportType)
	qs["filter[frequency]"] = string(f.Frequency)
	qs["filter[vendorNumber]"] = f.VendorNumber
	if !f.ReportDate.IsZero() {
		qs["filter[reportDate]"] = f.ReportDate.Format("2006-01-02")
	}
	if f.Version != "" {
		qs["filter[version]"] = string(f.Version)
	}
	return qs
}

//Validate sales report filter params
func (f *SalesReportsFilter) IsValid() error {
	if f.ReportType == "" {
		return fmt.Errorf("SalesReportsFilter@IsValid: %v", "ReportType is required")
	}
	if f.ReportSubType == "" {
		return fmt.Errorf("SalesReportsFilter@IsValid: %v", "ReportSubType is required")
	}
	if f.Frequency == "" {
		return fmt.Errorf("SalesReportsFilter@IsValid: %v", "Frequency is required")
	}
	return nil
}
