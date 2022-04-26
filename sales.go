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

//SalesReportSaleResponse struct
type SalesReportSaleResponse struct {
	*ResponseBody
	Data []*SalesReportSale `json:"data,omitempty"`
}

//GetReports Get sales report by filter
func (srr *SalesReportsResource) GetReports(ctx context.Context, filter SalesReportsFilterInterface) (*http.Response, error) {
	err := filter.IsValid()
	if err != nil {
		return nil, fmt.Errorf("SalesReportsResource.GetReports invalid filter: %v", err)
	}
	queryParams := srr.buildQueryParams(filter)
	return srr.transport.Get(ctx, "v1/salesReports", queryParams)
}

//GetSalesReportSale
func (srr *SalesReportsResource) GetSalesReportSale(ctx context.Context, filter *SalesReportsSalesFilter) (*SalesReportSaleResponse, *http.Response, error) {
	resp, err := srr.GetReports(ctx, filter)
	if err != nil {
		return nil, nil, fmt.Errorf("SalesReportsResource.GetSalesReportSale error: %v", err)
	}
	result := SalesReportSaleResponse{ResponseBody: &ResponseBody{}}
	result.status = resp.StatusCode
	if result.IsSuccess() {
		reports := []*SalesReportSale{}
		err = srr.unmarshalResponse(resp, &reports)
		if err != nil {
			return &result, resp, fmt.Errorf("SalesReportsResource.GetSalesReportSale error: %v", err)
		}
		result.Data = reports
	} else {
		err = srr.unmarshalResponse(resp, &result)
		if err != nil {
			return &result, resp, fmt.Errorf("SalesReportsResource.GetSalesReportSale error: %v", err)
		}
		return &result, resp, fmt.Errorf(result.GetError())
	}
	return &result, resp, nil
}

func (srr *SalesReportsResource) buildQueryParams(filter SalesReportsFilterInterface) map[string]interface{} {
	queryParams := filter.ToQueryParamsMap()
	queryParams["filter[vendorNumber]"] = srr.config.VendorNo
	return queryParams
}
