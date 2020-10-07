package appstore_sdk

import "time"

type SalesReportType string
type SalesReportSubType string
type SalesReportFrequency string
type SalesReportVersion string

const (
	FrequencyDaily   SalesReportFrequency = "DAILY"
	FrequencyWeekly  SalesReportFrequency = "WEEKLY"
	FrequencyMonthly SalesReportFrequency = "MONTHLY"
	FrequencyYearly  SalesReportFrequency = "YEARLY"
)

const (
	ReportTypeSales             SalesReportType = "SALES"
	ReportTypePreorder          SalesReportType = "PRE_ORDER"
	ReportTypeNewsStand         SalesReportType = "NEWSSTAND"
	ReportTypeSubscription      SalesReportType = "SUBSCRIPTION"
	ReportTypeSubscriptionEvent SalesReportType = "SUBSCRIPTION_EVENT"
	ReportTypeSubscriber        SalesReportType = "SUBSCRIBER"
)

const (
	ReportSubTypeSummary  SalesReportSubType = "SUMMARY"
	ReportSubTypeDetailed SalesReportSubType = "DETAILED"
	ReportSubTypeOptIn    SalesReportSubType = "OPT_IN"
)

const (
	Version10 SalesReportVersion = "1_0"
	Version12 SalesReportVersion = "1_2"
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

func (f *SalesReportsFilter) Detailed() *SalesReportsFilter {
	return f.SetReportSubType(ReportSubTypeDetailed)
}

func (f *SalesReportsFilter) OptIn() *SalesReportsFilter {
	return f.SetReportSubType(ReportSubTypeOptIn)
}

func (f *SalesReportsFilter) Summary() *SalesReportsFilter {
	return f.SetReportSubType(ReportSubTypeSummary)
}

func (f *SalesReportsFilter) SetReportType(value SalesReportType) *SalesReportsFilter {
	f.ReportType = value
	return f
}

func (f *SalesReportsFilter) Sales() *SalesReportsFilter {
	return f.SetReportType(ReportTypeSales)
}

func (f *SalesReportsFilter) PreOrder() *SalesReportsFilter {
	return f.SetReportType(ReportTypePreorder)
}

func (f *SalesReportsFilter) NewsStand() *SalesReportsFilter {
	return f.SetReportType(ReportTypeNewsStand)
}

func (f *SalesReportsFilter) Subscription() *SalesReportsFilter {
	return f.SetReportType(ReportTypeSubscription)
}

func (f *SalesReportsFilter) SubscriptionEvent() *SalesReportsFilter {
	return f.SetReportType(ReportTypeSubscriptionEvent)
}

func (f *SalesReportsFilter) Subscriber() *SalesReportsFilter {
	return f.SetReportType(ReportTypeSubscriber)
}

func (f *SalesReportsFilter) SetFrequency(value SalesReportFrequency) *SalesReportsFilter {
	f.Frequency = value
	return f
}

func (f *SalesReportsFilter) Daily() *SalesReportsFilter {
	return f.SetFrequency(FrequencyDaily)
}

func (f *SalesReportsFilter) Weekly() *SalesReportsFilter {
	return f.SetFrequency(FrequencyWeekly)
}

func (f *SalesReportsFilter) Monthly() *SalesReportsFilter {
	return f.SetFrequency(FrequencyMonthly)
}

func (f *SalesReportsFilter) Yearly() *SalesReportsFilter {
	return f.SetFrequency(FrequencyYearly)
}

func (f *SalesReportsFilter) SetVersion(value SalesReportVersion) *SalesReportsFilter {
	f.Version = value
	return f
}

func (f *SalesReportsFilter) Version12() *SalesReportsFilter {
	return f.SetVersion(Version12)
}

func (f *SalesReportsFilter) Version10() *SalesReportsFilter {
	return f.SetVersion(Version10)
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
