package appstore_sdk

import "fmt"

type SalesReportsResource struct {
	*ResourceAbstract
}

func (srr *SalesReportsResource) GetReport(filter *SalesReportsFilter) (*Response, error) {
	filter.VendorNumber = srr.config.VendorNo
	err := filter.IsValid()
	if err != nil {
		return nil, fmt.Errorf("SalesReportsResource@GetReport invalid filter: %v", err)
	}
	return srr.get("v1/salesReports", filter.ToQueryParamsMap())
}
