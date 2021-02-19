package appstore

import (
	"fmt"
	"time"
)

//SalesReportType type
type SalesReportType string

//SalesReportSubType type
type SalesReportSubType string

//SalesReportFrequency type
type SalesReportFrequency string

//SalesReportVersion type
type SalesReportVersion string

const (
	//SalesReportFrequencyDaily const
	SalesReportFrequencyDaily SalesReportFrequency = "DAILY"
	//SalesReportFrequencyWeekly const
	SalesReportFrequencyWeekly SalesReportFrequency = "WEEKLY"
	//SalesReportFrequencyMonthly const
	SalesReportFrequencyMonthly SalesReportFrequency = "MONTHLY"
	//SalesReportFrequencyYearly const
	SalesReportFrequencyYearly SalesReportFrequency = "YEARLY"
)

const (
	//SalesReportTypeSales const
	SalesReportTypeSales SalesReportType = "SALES"
	//SalesReportTypePreorder const
	SalesReportTypePreorder SalesReportType = "PRE_ORDER"
	//SalesReportTypeNewsStand const
	SalesReportTypeNewsStand SalesReportType = "NEWSSTAND"
	//SalesReportTypeSubscription const
	SalesReportTypeSubscription SalesReportType = "SUBSCRIPTION"
	//SalesReportTypeSubscriptionEvent const
	SalesReportTypeSubscriptionEvent SalesReportType = "SUBSCRIPTION_EVENT"
	//SalesReportTypeSubscriber const
	SalesReportTypeSubscriber SalesReportType = "SUBSCRIBER"
)

const (
	//SalesReportSubTypeSummary const
	SalesReportSubTypeSummary SalesReportSubType = "SUMMARY"
	//SalesReportSubTypeDetailed const
	SalesReportSubTypeDetailed SalesReportSubType = "DETAILED"
	//SalesReportSubTypeOptIn const
	SalesReportSubTypeOptIn SalesReportSubType = "OPT_IN"
)

const (
	//SalesReportVersion10 const
	SalesReportVersion10 SalesReportVersion = "1_0"
	//SalesReportVersion12 const
	SalesReportVersion12 SalesReportVersion = "1_2"
)

//SalesReportsFilter Sales reports filter
type SalesReportsFilter struct {
	ReportDate    time.Time            //The report date to download. The date is specified in the YYYY-MM-DD format for all report frequencies except DAILY, which does not require a specified date. For more information, see report availability and storage.
	ReportSubType SalesReportSubType   //(Required) The report sub type to download. For a list of values, see Allowed Values Based on Sales Report Type below. Possible values: SUMMARY, DETAILED, OPT_IN
	ReportType    SalesReportType      //(Required) The report to download. For more details on each report type see About Reports. Possible values: SALES, PRE_ORDER, NEWSSTAND, SUBSCRIPTION, SUBSCRIPTION_EVENT, SUBSCRIBER
	Frequency     SalesReportFrequency //(Required) Frequency of the report to download. For a list of values, see Allowed Values Based on Sales Report Type below. Possible values: DAILY, WEEKLY, MONTHLY, YEARLY
	Version       SalesReportVersion   //The version of the report. For a list of values, see Allowed Values Based on Sales Report Type below.
	VendorNumber  string               //(Required) You can find your vendor number in Payments and Financial Reports.
}

//SetReportDate Set report date
func (f *SalesReportsFilter) SetReportDate(value time.Time) *SalesReportsFilter {
	f.ReportDate = value
	return f
}

//SetReportSubType Set report sub type
func (f *SalesReportsFilter) SetReportSubType(value SalesReportSubType) *SalesReportsFilter {
	f.ReportSubType = value
	return f
}

//SubTypeDetailed Change report sub type to Detailed
func (f *SalesReportsFilter) SubTypeDetailed() *SalesReportsFilter {
	return f.SetReportSubType(SalesReportSubTypeDetailed)
}

//SubTypeOptIn Change report sub type to OptIn
func (f *SalesReportsFilter) SubTypeOptIn() *SalesReportsFilter {
	return f.SetReportSubType(SalesReportSubTypeOptIn)
}

//SubTypeSummary Change report sub type to Summary
func (f *SalesReportsFilter) SubTypeSummary() *SalesReportsFilter {
	return f.SetReportSubType(SalesReportSubTypeSummary)
}

//SetReportType Set report type
func (f *SalesReportsFilter) SetReportType(value SalesReportType) *SalesReportsFilter {
	f.ReportType = value
	return f
}

//TypeSales Change report type to Sales
func (f *SalesReportsFilter) TypeSales() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypeSales)
}

//TypePreOrder Change report type to PreOrder
func (f *SalesReportsFilter) TypePreOrder() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypePreorder)
}

//TypeNewsStand Change report type to NewsStand
func (f *SalesReportsFilter) TypeNewsStand() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypeNewsStand)
}

//TypeSubscription Change report type to Subscription
func (f *SalesReportsFilter) TypeSubscription() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypeSubscription)
}

//TypeSubscriptionEvent Change report type to SubscriptionEvent
func (f *SalesReportsFilter) TypeSubscriptionEvent() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypeSubscriptionEvent)
}

//TypeSubscriber Change report type to Subscriber
func (f *SalesReportsFilter) TypeSubscriber() *SalesReportsFilter {
	return f.SetReportType(SalesReportTypeSubscriber)
}

//SetFrequency Set frequency
func (f *SalesReportsFilter) SetFrequency(value SalesReportFrequency) *SalesReportsFilter {
	f.Frequency = value
	return f
}

//Daily Change frequency to Daily
func (f *SalesReportsFilter) Daily() *SalesReportsFilter {
	return f.SetFrequency(SalesReportFrequencyDaily)
}

//Weekly Change frequency to Weekly
func (f *SalesReportsFilter) Weekly() *SalesReportsFilter {
	return f.SetFrequency(SalesReportFrequencyWeekly)
}

//Monthly Change frequency to Monthly
func (f *SalesReportsFilter) Monthly() *SalesReportsFilter {
	return f.SetFrequency(SalesReportFrequencyMonthly)
}

//Yearly Change frequency to Yearly
func (f *SalesReportsFilter) Yearly() *SalesReportsFilter {
	return f.SetFrequency(SalesReportFrequencyYearly)
}

//SetVersion Set version
func (f *SalesReportsFilter) SetVersion(value SalesReportVersion) *SalesReportsFilter {
	f.Version = value
	return f
}

//Version12 Change version to 1_2
func (f *SalesReportsFilter) Version12() *SalesReportsFilter {
	return f.SetVersion(SalesReportVersion12)
}

//Version10 Change version to 1_0
func (f *SalesReportsFilter) Version10() *SalesReportsFilter {
	return f.SetVersion(SalesReportVersion10)
}

//ToQueryParamsMap Convert filter to query params
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

//IsValid Validate sales report filter params
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
