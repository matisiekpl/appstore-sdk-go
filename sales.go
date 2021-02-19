package appstore

import "fmt"

//Sales reports resource
type SalesReportsResource struct {
	*ResourceAbstract
}

//Get sales report by filter
//@unmarshal SalesReportSale, SalesReportSubscription, SalesReportSubscriptionEvent, SalesReportSubscriber
func (srr *SalesReportsResource) GetReport(filter *SalesReportsFilter) (*Response, error) {
	filter.VendorNumber = srr.config.VendorNo
	err := filter.IsValid()
	if err != nil {
		return nil, fmt.Errorf("SalesReportsResource@GetReport invalid filter: %v", err)
	}
	return srr.get("v1/salesReports", filter.ToQueryParamsMap())
}
