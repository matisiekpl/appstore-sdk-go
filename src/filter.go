package appstore_sdk

import "time"

type ReportType string
type ReportSubType string
type Frequency string
type Version string

const (
	DAILY   Frequency = "DAILY"
	WEEKLY  Frequency = "WEEKLY"
	MONTHLY Frequency = "MONTHLY"
	YEARLY  Frequency = "YEARLY"
)

const (
	SALES              ReportType = "SALES"
	PRE_ORDER          ReportType = "PRE_ORDER"
	NEWSSTAND          ReportType = "NEWSSTAND"
	SUBSCRIPTION       ReportType = "SUBSCRIPTION"
	SUBSCRIPTION_EVENT ReportType = "SUBSCRIPTION_EVENT"
	SUBSCRIBER         ReportType = "SUBSCRIBER"
)

const (
	SUMMARY  ReportSubType = "SUMMARY"
	DETAILED ReportSubType = "DETAILED"
	OPT_IN   ReportSubType = "OPT_IN"
)

const (
	V1_0 Version = "1_0"
	V1_2 Version = "1_2"
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
	return f.SetReportSubType(DETAILED)
}

func (f *SalesReportsFilter) OptIn() *SalesReportsFilter {
	return f.SetReportSubType(OPT_IN)
}

func (f *SalesReportsFilter) Summary() *SalesReportsFilter {
	return f.SetReportSubType(SUMMARY)
}

func (f *SalesReportsFilter) SetReportType(value ReportType) *SalesReportsFilter {
	f.ReportType = value
	return f
}

func (f *SalesReportsFilter) Sales() *SalesReportsFilter {
	return f.SetReportType(SALES)
}

func (f *SalesReportsFilter) PreOrder() *SalesReportsFilter {
	return f.SetReportType(PRE_ORDER)
}

func (f *SalesReportsFilter) NewsStand() *SalesReportsFilter {
	return f.SetReportType(NEWSSTAND)
}

func (f *SalesReportsFilter) Subscription() *SalesReportsFilter {
	return f.SetReportType(SUBSCRIPTION)
}

func (f *SalesReportsFilter) SubscriptionEvent() *SalesReportsFilter {
	return f.SetReportType(SUBSCRIPTION_EVENT)
}

func (f *SalesReportsFilter) Subscriber() *SalesReportsFilter {
	return f.SetReportType(SUBSCRIBER)
}

func (f *SalesReportsFilter) SetFrequency(value Frequency) *SalesReportsFilter {
	f.Frequency = value
	return f
}

func (f *SalesReportsFilter) Daily() *SalesReportsFilter {
	return f.SetFrequency(DAILY)
}

func (f *SalesReportsFilter) Weekly() *SalesReportsFilter {
	return f.SetFrequency(WEEKLY)
}

func (f *SalesReportsFilter) Monthly() *SalesReportsFilter {
	return f.SetFrequency(MONTHLY)
}

func (f *SalesReportsFilter) Yearly() *SalesReportsFilter {
	return f.SetFrequency(YEARLY)
}

func (f *SalesReportsFilter) SetVersion(value Version) *SalesReportsFilter {
	f.Version = value
	return f
}

func (f *SalesReportsFilter) V1_2() *SalesReportsFilter {
	return f.SetVersion(V1_2)
}

func (f *SalesReportsFilter) V1_0() *SalesReportsFilter {
	return f.SetVersion(V1_0)
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
