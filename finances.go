package appstore

import (
	"context"
	"fmt"
	"net/http"
)

//FinancialReportsResponse struct
type FinancialReportsResponse struct {
	ResponseBody
	Data []*FinancialReport `json:"data,omitempty"`
}

//FinancesReportsResource reports
type FinancesReportsResource struct {
	ResourceAbstract
}

//GetReports Get finances reports by filter
func (frr *FinancesReportsResource) GetReports(ctx context.Context, filter *FinancesReportsFilter) (*http.Response, error) {
	err := filter.IsValid()
	if err != nil {
		return nil, fmt.Errorf("FinancesReportsResource.GetReports invalid filter: %v", err)
	}
	queryParams := frr.buildQueryParams(filter)
	return frr.transport.Get(ctx, "v1/financeReports", queryParams)
}

//GetFinancialReports
func (frr *FinancesReportsResource) GetFinancialReports(ctx context.Context, filter *FinancesReportsFilter) (*FinancialReportsResponse, *http.Response, error) {
	resp, err := frr.GetReports(ctx, filter)
	if err != nil {
		return nil, nil, fmt.Errorf("FinancialReportsResponse.GetFinancialReports error: %v", err)
	}
	result := FinancialReportsResponse{}
	result.status = resp.StatusCode
	if result.IsSuccess() {
		reports := []*FinancialReport{}
		err = frr.unmarshalResponse(resp, &reports, true)
		if err != nil {
			return &result, resp, fmt.Errorf("FinancesReportsResource.GetFinancialReports error: %v", err)
		}
		result.Data = reports
	} else {
		err = frr.unmarshalResponse(resp, &result, false)
		if err != nil {
			return &result, resp, fmt.Errorf("FinancesReportsResource.GetFinancialReports error: %v", err)
		}
		return &result, resp, fmt.Errorf(result.GetError())
	}
	return &result, resp, nil
}

//buildQueryParams
func (frr *FinancesReportsResource) buildQueryParams(filter *FinancesReportsFilter) map[string]interface{} {
	queryParams := filter.toQueryParamsMap()
	queryParams["filter[vendorNumber]"] = frr.config.VendorNo
	return queryParams
}
