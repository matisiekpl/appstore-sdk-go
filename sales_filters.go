package appstore

import (
	"fmt"
	"time"
)

// SalesReportType type
type SalesReportType string

// SalesReportSubType type
type SalesReportSubType string

// SalesReportFrequency type
type SalesReportFrequency string

// SalesReportVersion type
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
	//SalesReportTypeSubscriptionOfferCodeRedemption const
	SalesReportTypeSubscriptionOfferCodeRedemption SalesReportType = "SUBSCRIPTION_OFFER_CODE_REDEMPTION"
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
	//SalesReportVersion11 const
	SalesReportVersion11 SalesReportVersion = "1_1"
	//SalesReportVersion12 const
	SalesReportVersion12 SalesReportVersion = "1_2"
	//SalesReportVersion13 const
	SalesReportVersion13 SalesReportVersion = "1_3"
)

// SalesReportsBaseFilter Sales reports filter
type SalesReportsBaseFilter struct {
	ReportDate    time.Time            //The report date to download. The date is specified in the YYYY-MM-DD format for all report frequencies except DAILY, which does not require a specified date. For more information, see report availability and storage.
	ReportSubType SalesReportSubType   //(Required) The report sub type to download. For a list of values, see Allowed Values Based on Sales Report Type below. Possible values: SUMMARY, DETAILED, OPT_IN
	ReportType    SalesReportType      //(Required) The report to download. For more details on each report type see About Reports. Possible values: SALES, PRE_ORDER, NEWSSTAND, SUBSCRIPTION, SUBSCRIPTION_EVENT, SUBSCRIBER
	Frequency     SalesReportFrequency //(Required) Frequency of the report to download. For a list of values, see Allowed Values Based on Sales Report Type below. Possible values: DAILY, WEEKLY, MONTHLY, YEARLY
	Version       SalesReportVersion   //The version of the report. For a list of values, see Allowed Values Based on Sales Report Type below.
}

// SetReportDate Set report date
func (f *SalesReportsBaseFilter) SetReportDate(value time.Time) *SalesReportsBaseFilter {
	f.ReportDate = value
	return f
}

// SetReportSubType Set report sub type
func (f *SalesReportsBaseFilter) SetReportSubType(value SalesReportSubType) *SalesReportsBaseFilter {
	f.ReportSubType = value
	return f
}

// SubTypeDetailed Change report sub type to Detailed
func (f *SalesReportsBaseFilter) SubTypeDetailed() *SalesReportsBaseFilter {
	return f.SetReportSubType(SalesReportSubTypeDetailed)
}

// SubTypeOptIn Change report sub type to OptIn
func (f *SalesReportsBaseFilter) SubTypeOptIn() *SalesReportsBaseFilter {
	return f.SetReportSubType(SalesReportSubTypeOptIn)
}

// SubTypeSummary Change report sub type to Summary
func (f *SalesReportsBaseFilter) SubTypeSummary() *SalesReportsBaseFilter {
	return f.SetReportSubType(SalesReportSubTypeSummary)
}

// SetReportType Set report type
func (f *SalesReportsBaseFilter) SetReportType(value SalesReportType) *SalesReportsBaseFilter {
	f.ReportType = value
	return f
}

// TypeSales Change report type to Sales
func (f *SalesReportsBaseFilter) TypeSales() *SalesReportsBaseFilter {
	return f.SetReportType(SalesReportTypeSales)
}

// TypePreOrder Change report type to PreOrder
func (f *SalesReportsBaseFilter) TypePreOrder() *SalesReportsBaseFilter {
	return f.SetReportType(SalesReportTypePreorder)
}

// TypeNewsStand Change report type to NewsStand
func (f *SalesReportsBaseFilter) TypeNewsStand() *SalesReportsBaseFilter {
	return f.SetReportType(SalesReportTypeNewsStand)
}

// TypeSubscription Change report type to Subscription
func (f *SalesReportsBaseFilter) TypeSubscription() *SalesReportsBaseFilter {
	return f.SetReportType(SalesReportTypeSubscription)
}

// TypeSubscriptionEvent Change report type to SubscriptionEvent
func (f *SalesReportsBaseFilter) TypeSubscriptionEvent() *SalesReportsBaseFilter {
	return f.SetReportType(SalesReportTypeSubscriptionEvent)
}

// TypeSubscriber Change report type to Subscriber
func (f *SalesReportsBaseFilter) TypeSubscriber() *SalesReportsBaseFilter {
	return f.SetReportType(SalesReportTypeSubscriber)
}

// SetFrequency Set frequency
func (f *SalesReportsBaseFilter) SetFrequency(value SalesReportFrequency) *SalesReportsBaseFilter {
	f.Frequency = value
	return f
}

// Daily Change frequency to Daily
func (f *SalesReportsBaseFilter) Daily() *SalesReportsBaseFilter {
	return f.SetFrequency(SalesReportFrequencyDaily)
}

// Weekly Change frequency to Weekly
func (f *SalesReportsBaseFilter) Weekly() *SalesReportsBaseFilter {
	return f.SetFrequency(SalesReportFrequencyWeekly)
}

// Monthly Change frequency to Monthly
func (f *SalesReportsBaseFilter) Monthly() *SalesReportsBaseFilter {
	return f.SetFrequency(SalesReportFrequencyMonthly)
}

// Yearly Change frequency to Yearly
func (f *SalesReportsBaseFilter) Yearly() *SalesReportsBaseFilter {
	return f.SetFrequency(SalesReportFrequencyYearly)
}

// SetVersion Set version
func (f *SalesReportsBaseFilter) SetVersion(value SalesReportVersion) *SalesReportsBaseFilter {
	f.Version = value
	return f
}

// Version13 Change version to 1_3
func (f *SalesReportsBaseFilter) Version13() *SalesReportsBaseFilter {
	return f.SetVersion(SalesReportVersion13)
}

// Version12 Change version to 1_2
func (f *SalesReportsBaseFilter) Version12() *SalesReportsBaseFilter {
	return f.SetVersion(SalesReportVersion12)
}

// Version11 Change version to 1_1
func (f *SalesReportsBaseFilter) Version11() *SalesReportsBaseFilter {
	return f.SetVersion(SalesReportVersion11)
}

// Version10 Change version to 1_0
func (f *SalesReportsBaseFilter) Version10() *SalesReportsBaseFilter {
	return f.SetVersion(SalesReportVersion10)
}

// ToQueryParamsMap Convert filter to query params
func (f *SalesReportsBaseFilter) ToQueryParamsMap() map[string]interface{} {
	qs := make(map[string]interface{})
	qs["filter[reportSubType]"] = string(f.ReportSubType)
	qs["filter[reportType]"] = string(f.ReportType)
	qs["filter[frequency]"] = string(f.Frequency)
	if !f.ReportDate.IsZero() {
		qs["filter[reportDate]"] = f.ReportDate.Format("2006")
	}
	if f.Version != "" {
		qs["filter[version]"] = string(f.Version)
	}
	return qs
}

// IsValid Validate sales report filter params
func (f *SalesReportsBaseFilter) IsValid() error {
	if f.ReportType == "" {
		return fmt.Errorf("SalesReportsBaseFilter.IsValid: %v", "ReportType is required")
	}
	if f.ReportSubType == "" {
		return fmt.Errorf("SalesReportsBaseFilter.IsValid: %v", "ReportSubType is required")
	}
	if f.Frequency == "" {
		return fmt.Errorf("SalesReportsBaseFilter.IsValid: %v", "Frequency is required")
	}
	return nil
}

// SalesReportsFilterInterface
type SalesReportsFilterInterface interface {
	IsValid() error
	ToQueryParamsMap() map[string]interface{}
}

// SalesReportsFilter
type SalesReportsFilter struct {
	SalesReportsBaseFilter
}

// IsValid Validate sales report filter params
func (f *SalesReportsFilter) IsValid() error {
	err := f.SalesReportsBaseFilter.IsValid()
	if err != nil {
		return err
	}
	if f.ReportType != SalesReportTypeSales {
		return fmt.Errorf("SalesReportsFilter.IsValid: %v", "ReportType is not valid")
	}
	if f.ReportSubType != SalesReportSubTypeSummary {
		return fmt.Errorf("SalesReportsFilter.IsValid: %v", "ReportSubType is not valid")
	}
	if f.Version != SalesReportVersion11 && f.Version != SalesReportVersion10 {
		return fmt.Errorf("SalesReportsFilter.IsValid: %v", "Version is not valid")
	}
	return nil
}

// SubscriptionsReportsFilter
type SubscriptionsReportsFilter struct {
	SalesReportsBaseFilter
}

// IsValid Validate sales report filter params
func (f *SubscriptionsReportsFilter) IsValid() error {
	err := f.SalesReportsBaseFilter.IsValid()
	if err != nil {
		return err
	}
	if f.ReportType != SalesReportTypeSubscription {
		return fmt.Errorf("SubscriptionsReportsFilter.IsValid: %v", "ReportType is not valid")
	}
	if f.ReportSubType != SalesReportSubTypeSummary {
		return fmt.Errorf("SubscriptionsReportsFilter.IsValid: %v", "ReportSubType is not valid")
	}
	if f.Frequency != SalesReportFrequencyDaily {
		return fmt.Errorf("SubscriptionsReportsFilter.IsValid: %v", "Frequency is not valid")
	}
	if f.Version != SalesReportVersion12 && f.Version != SalesReportVersion13 {
		return fmt.Errorf("SubscriptionsReportsFilter.IsValid: %v", "Version is not valid")
	}
	return nil
}

// SubscriptionsEventsReportsFilter
type SubscriptionsEventsReportsFilter struct {
	SalesReportsBaseFilter
}

// IsValid Validate sales report filter params
func (f *SubscriptionsEventsReportsFilter) IsValid() error {
	err := f.SalesReportsBaseFilter.IsValid()
	if err != nil {
		return err
	}
	if f.ReportType != SalesReportTypeSubscriptionEvent {
		return fmt.Errorf("SubscriptionsEventsReportsFilter.IsValid: %v", "ReportType is not valid")
	}
	if f.ReportSubType != SalesReportSubTypeSummary {
		return fmt.Errorf("SubscriptionsEventsReportsFilter.IsValid: %v", "ReportSubType is not valid")
	}
	if f.Frequency != SalesReportFrequencyDaily {
		return fmt.Errorf("SubscriptionsEventsReportsFilter.IsValid: %v", "Frequency is not valid")
	}
	if f.Version != SalesReportVersion12 && f.Version != SalesReportVersion13 {
		return fmt.Errorf("SubscriptionsEventsReportsFilter.IsValid: %v", "Version is not valid")
	}
	return nil
}

// SubscribersReportsFilter
type SubscribersReportsFilter struct {
	SalesReportsBaseFilter
}

// IsValid Validate sales report filter params
func (f *SubscribersReportsFilter) IsValid() error {
	err := f.SalesReportsBaseFilter.IsValid()
	if err != nil {
		return err
	}
	if f.ReportType != SalesReportTypeSubscriber {
		return fmt.Errorf("SubscribersReportsFilter.IsValid: %v", "ReportType is not valid")
	}
	if f.ReportSubType != SalesReportSubTypeDetailed {
		return fmt.Errorf("SubscribersReportsFilter.IsValid: %v", "ReportSubType is not valid")
	}
	if f.Frequency != SalesReportFrequencyDaily {
		return fmt.Errorf("SubscribersReportsFilter.IsValid: %v", "Frequency is not valid")
	}
	if f.Version != SalesReportVersion12 && f.Version != SalesReportVersion13 {
		return fmt.Errorf("SubscribersReportsFilter.IsValid: %v", "Version is not valid")
	}
	return nil
}

// SubscriptionsOffersCodesRedemptionReportsFilter
type SubscriptionsOffersCodesRedemptionReportsFilter struct {
	SalesReportsBaseFilter
}

// IsValid Validate sales report filter params
func (f *SubscriptionsOffersCodesRedemptionReportsFilter) IsValid() error {
	err := f.SalesReportsBaseFilter.IsValid()
	if err != nil {
		return err
	}
	if f.ReportType != SalesReportTypeSubscriptionOfferCodeRedemption {
		return fmt.Errorf("SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: %v", "ReportType is not valid")
	}
	if f.ReportSubType != SalesReportSubTypeSummary {
		return fmt.Errorf("SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: %v", "ReportSubType is not valid")
	}
	if f.Frequency != SalesReportFrequencyDaily {
		return fmt.Errorf("SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: %v", "Frequency is not valid")
	}
	if f.Version != SalesReportVersion10 {
		return fmt.Errorf("SubscriptionsOffersCodesRedemptionReportsFilter.IsValid: %v", "Version is not valid")
	}
	return nil
}

// NewsstandReportsFilter
type NewsstandReportsFilter struct {
	SalesReportsBaseFilter
}

// IsValid Validate sales report filter params
func (f *NewsstandReportsFilter) IsValid() error {
	err := f.SalesReportsBaseFilter.IsValid()
	if err != nil {
		return err
	}
	if f.ReportType != SalesReportTypeNewsStand {
		return fmt.Errorf("NewsstandReportsFilter.IsValid: %v", "ReportType is not valid")
	}
	if f.ReportSubType != SalesReportSubTypeDetailed {
		return fmt.Errorf("NewsstandReportsFilter.IsValid: %v", "ReportSubType is not valid")
	}
	if f.Frequency != SalesReportFrequencyDaily && f.Frequency != SalesReportFrequencyWeekly {
		return fmt.Errorf("NewsstandReportsFilter.IsValid: %v", "Frequency is not valid")
	}
	if f.Version != SalesReportVersion10 {
		return fmt.Errorf("NewsstandReportsFilter.IsValid: %v", "Version is not valid")
	}
	return nil
}

// PreOrdersReportsFilter
type PreOrdersReportsFilter struct {
	SalesReportsBaseFilter
}

// IsValid Validate sales report filter params
func (f *PreOrdersReportsFilter) IsValid() error {
	err := f.SalesReportsBaseFilter.IsValid()
	if err != nil {
		return err
	}
	if f.ReportType != SalesReportTypePreorder {
		return fmt.Errorf("PreOrdersReportsFilter.IsValid: %v", "ReportType is not valid")
	}
	if f.ReportSubType != SalesReportSubTypeSummary {
		return fmt.Errorf("PreOrdersReportsFilter.IsValid: %v", "ReportSubType is not valid")
	}
	if f.Version != SalesReportVersion10 {
		return fmt.Errorf("PreOrdersReportsFilter.IsValid: %v", "Version is not valid")
	}
	return nil
}

// NewSalesReportsFilter
func NewSalesReportsFilter() *SalesReportsFilter {
	return &SalesReportsFilter{SalesReportsBaseFilter{ReportType: SalesReportTypeSales}}
}

// NewSubscriptionsReportsFilter
func NewSubscriptionsReportsFilter() *SubscriptionsReportsFilter {
	return &SubscriptionsReportsFilter{SalesReportsBaseFilter{ReportType: SalesReportTypeSubscription}}
}

// NewSubscriptionsEventsReportsFilter
func NewSubscriptionsEventsReportsFilter() *SubscriptionsEventsReportsFilter {
	return &SubscriptionsEventsReportsFilter{SalesReportsBaseFilter{ReportType: SalesReportTypeSubscriptionEvent}}
}

// NewSubscribersReportsFilter
func NewSubscribersReportsFilter() *SubscribersReportsFilter {
	return &SubscribersReportsFilter{SalesReportsBaseFilter{ReportType: SalesReportTypeSubscriber}}
}

func NewSubscriptionsOffersCodesRedemptionReportsFilter() *SubscriptionsOffersCodesRedemptionReportsFilter {
	return &SubscriptionsOffersCodesRedemptionReportsFilter{SalesReportsBaseFilter{ReportType: SalesReportTypeSubscriptionOfferCodeRedemption}}
}

// NewNewsstandReportsFilter
func NewNewsstandReportsFilter() *NewsstandReportsFilter {
	return &NewsstandReportsFilter{SalesReportsBaseFilter{ReportType: SalesReportTypeNewsStand}}
}

// NewPreOrdersReportsFilter
func NewPreOrdersReportsFilter() *PreOrdersReportsFilter {
	return &PreOrdersReportsFilter{SalesReportsBaseFilter{ReportType: SalesReportTypePreorder}}
}
