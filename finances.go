package appstore

import (
	"context"
	"fmt"
	"net/http"
)

//FinancesReportsResource reports
type FinancesReportsResource struct {
	*ResourceAbstract
}

//GetReports Get financial reports by filter
func (frr *FinancesReportsResource) GetReports(ctx context.Context, filter FinancesReportsFilter) (*http.Response, error) {
	err := filter.IsValid()
	if err != nil {
		return nil, fmt.Errorf("FinancesReportsResource.GetReports invalid filter: %v", err)
	}
	queryParams := frr.buildQueryParams(filter)
	return frr.transport.Get(ctx, "v1/financeReports", queryParams)
}

//buildQueryParams
func (frr *FinancesReportsResource) buildQueryParams(filter FinancesReportsFilter) map[string]interface{} {
	queryParams := filter.ToQueryParamsMap()
	queryParams["filter[vendorNumber]"] = frr.config.VendorNo
	return queryParams
}
