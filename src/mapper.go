package appstore_sdk

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"io"
)

type Mapper struct {
}

func (rp *Mapper) Init() {
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.LazyQuotes = true
		r.Comma = '\t'
		return r
	})
}

func (rp *Mapper) MapReportsSalesResponse(response []byte) ([]*SalesReportSale, error) {
	reports := []*SalesReportSale{}
	err := gocsv.UnmarshalBytes(response, &reports)
	return reports, err
}

func (rp *Mapper) MapReportsSubscribersResponse(response []byte) ([]*SalesReportSubscriber, error) {
	reports := []*SalesReportSubscriber{}
	err := gocsv.UnmarshalBytes(response, &reports)
	return reports, err
}

func (rp *Mapper) MapReportsSubscriptionsResponse(response []byte) ([]*SalesReportSubscription, error) {
	reports := []*SalesReportSubscription{}
	err := gocsv.UnmarshalBytes(response, &reports)
	return reports, err
}

func (rp *Mapper) MapReportsSubscriptionsEventsResponse(response []byte) ([]*SalesReportSubscriptionEvent, error) {
	reports := []*SalesReportSubscriptionEvent{}
	err := gocsv.UnmarshalBytes(response, &reports)
	return reports, err
}
