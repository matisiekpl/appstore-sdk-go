package appstore

import (
	"context"
	"fmt"
	"net/http"
)

//SalesReportsResource reports
type SalesReportsResource struct {
	*ResourceAbstract
}

//GetReport Get sales report by filter
func (srr *SalesReportsResource) GetReport(ctx context.Context, filter *SalesReportsFilter) (*http.Response, error) {
	filter.VendorNumber = srr.config.VendorNo
	err := filter.IsValid()
	if err != nil {
		return nil, fmt.Errorf("SalesReportsResource@GetReport invalid filter: %v", err)
	}
	return srr.transport.Get(ctx, "v1/salesReports", filter.ToQueryParamsMap())
}
