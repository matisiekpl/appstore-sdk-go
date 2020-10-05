package appstore_sdk

import "time"

type ReportType string
type ReportSubType string
type Frequency string
type Version string

const (
	FrequencyDaily   Frequency = "DAILY"
	FrequencyWeekly  Frequency = "WEEKLY"
	FrequencyMonthly Frequency = "MONTHLY"
	FrequencyYearly  Frequency = "YEARLY"
)

const (
	ReportTypeSales             ReportType = "SALES"
	ReportTypePreorder          ReportType = "PRE_ORDER"
	ReportTypeNewsStand         ReportType = "NEWSSTAND"
	ReportTypeSubscription      ReportType = "SUBSCRIPTION"
	ReportTypeSubscriptionEvent ReportType = "SUBSCRIPTION_EVENT"
	ReportTypeSubscriber        ReportType = "SUBSCRIBER"
)

const (
	ReportSubTypeSummary  ReportSubType = "SUMMARY"
	ReportSubTypeDetailed ReportSubType = "DETAILED"
	ReportSubTypeOptIn    ReportSubType = "OPT_IN"
)

const (
	Version10 Version = "1_0"
	Version12 Version = "1_2"
)

type SalesReportsFilter struct {
	ReportDate    time.Time
	ReportSubType ReportSubType
	ReportType    ReportType
	Frequency     Frequency
	Version       Version
	VendorNumber  string
}

func (f *SalesReportsFilter) SetReportDate(value time.Time) *SalesReportsFilter {
	f.ReportDate = value
	return f
}

func (f *SalesReportsFilter) SetReportSubType(value ReportSubType) *SalesReportsFilter {
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

func (f *SalesReportsFilter) SetReportType(value ReportType) *SalesReportsFilter {
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

func (f *SalesReportsFilter) SetFrequency(value Frequency) *SalesReportsFilter {
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

func (f *SalesReportsFilter) SetVersion(value Version) *SalesReportsFilter {
	f.Version = value
	return f
}

func (f *SalesReportsFilter) Version12() *SalesReportsFilter {
	return f.SetVersion(Version12)
}

func (f *SalesReportsFilter) Version10() *SalesReportsFilter {
	return f.SetVersion(Version10)
}

func (f *SalesReportsFilter) ToQueryParamsMap() map[string]string {
	qs := make(map[string]string)
	qs["filter[reportDate]"] = f.ReportDate.Format("2006-01-02")
	qs["filter[reportSubType]"] = string(f.ReportSubType)
	qs["filter[reportType]"] = string(f.ReportType)
	qs["filter[frequency]"] = string(f.Frequency)
	qs["filter[version]"] = string(f.Version)
	qs["filter[vendorNumber]"] = f.VendorNumber
	return qs
}
