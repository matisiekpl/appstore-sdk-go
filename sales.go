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

//SalesReportsResponse struct
type SalesReportsResponse struct {
	*ResponseBody
	Data []*SalesReport `json:"data,omitempty"`
}

//SubscriptionsReportsResponse struct
type SubscriptionsReportsResponse struct {
	*ResponseBody
	Data []*SubscriptionsReport `json:"data,omitempty"`
}

//SubscriptionsEventsReportsResponse struct
type SubscriptionsEventsReportsResponse struct {
	*ResponseBody
	Data []*SubscriptionsEventsReport `json:"data,omitempty"`
}

//SubscribersReportsResponse struct
type SubscribersReportsResponse struct {
	*ResponseBody
	Data []*SubscribersReport `json:"data,omitempty"`
}

//PreOrdersReportsResponse struct
type PreOrdersReportsResponse struct {
	*ResponseBody
	Data []*PreOrdersReport `json:"data,omitempty"`
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

//GetSalesReports
func (srr *SalesReportsResource) GetSalesReports(ctx context.Context, filter *SalesReportsFilter) (*SalesReportsResponse, *http.Response, error) {
	resp, err := srr.GetReports(ctx, filter)
	if err != nil {
		return nil, nil, fmt.Errorf("SalesReportsResource.GetSalesReports error: %v", err)
	}
	result := SalesReportsResponse{ResponseBody: &ResponseBody{}}
	result.status = resp.StatusCode
	if result.IsSuccess() {
		reports := []*SalesReport{}
		err = srr.unmarshalResponse(resp, &reports)
		if err != nil {
			return &result, resp, fmt.Errorf("SalesReportsResource.GetSalesReports error: %v", err)
		}
		result.Data = reports
	} else {
		err = srr.unmarshalResponse(resp, &result)
		if err != nil {
			return &result, resp, fmt.Errorf("SalesReportsResource.GetSalesReports error: %v", err)
		}
		return &result, resp, fmt.Errorf(result.GetError())
	}
	return &result, resp, nil
}

//GetSubscriptionsReports
func (srr *SalesReportsResource) GetSubscriptionsReports(ctx context.Context, filter *SubscriptionsReportsFilter) (*SubscriptionsReportsResponse, *http.Response, error) {
	resp, err := srr.GetReports(ctx, filter)
	if err != nil {
		return nil, nil, fmt.Errorf("SalesReportsResource.GetSubscriptionsReports error: %v", err)
	}
	result := SubscriptionsReportsResponse{ResponseBody: &ResponseBody{}}
	result.status = resp.StatusCode
	if result.IsSuccess() {
		reports := []*SubscriptionsReport{}
		err = srr.unmarshalResponse(resp, &reports)
		if err != nil {
			return &result, resp, fmt.Errorf("SalesReportsResource.GetSubscriptionsReports error: %v", err)
		}
		result.Data = reports
	} else {
		err = srr.unmarshalResponse(resp, &result)
		if err != nil {
			return &result, resp, fmt.Errorf("SalesReportsResource.GetSubscriptionsReports error: %v", err)
		}
		return &result, resp, fmt.Errorf(result.GetError())
	}
	return &result, resp, nil
}

//GetSubscriptionsEventsReports
func (srr *SalesReportsResource) GetSubscriptionsEventsReports(ctx context.Context, filter *SubscriptionsEventsReportsFilter) (*SubscriptionsEventsReportsResponse, *http.Response, error) {
	resp, err := srr.GetReports(ctx, filter)
	if err != nil {
		return nil, nil, fmt.Errorf("SalesReportsResource.GetSubscriptionsReports error: %v", err)
	}
	result := SubscriptionsEventsReportsResponse{ResponseBody: &ResponseBody{}}
	result.status = resp.StatusCode
	if result.IsSuccess() {
		reports := []*SubscriptionsEventsReport{}
		err = srr.unmarshalResponse(resp, &reports)
		if err != nil {
			return &result, resp, fmt.Errorf("SalesReportsResource.GetSubscriptionsReports error: %v", err)
		}
		result.Data = reports
	} else {
		err = srr.unmarshalResponse(resp, &result)
		if err != nil {
			return &result, resp, fmt.Errorf("SalesReportsResource.GetSubscriptionsReports error: %v", err)
		}
		return &result, resp, fmt.Errorf(result.GetError())
	}
	return &result, resp, nil
}

//GetSubscriptionsReports
func (srr *SalesReportsResource) GetSubscribersReports(ctx context.Context, filter *SubscribersReportsFilter) (*SubscribersReportsResponse, *http.Response, error) {
	resp, err := srr.GetReports(ctx, filter)
	if err != nil {
		return nil, nil, fmt.Errorf("SalesReportsResource.GetSubscribersReports error: %v", err)
	}
	result := SubscribersReportsResponse{ResponseBody: &ResponseBody{}}
	result.status = resp.StatusCode
	if result.IsSuccess() {
		reports := []*SubscribersReport{}
		err = srr.unmarshalResponse(resp, &reports)
		if err != nil {
			return &result, resp, fmt.Errorf("SalesReportsResource.GetSubscribersReports error: %v", err)
		}
		result.Data = reports
	} else {
		err = srr.unmarshalResponse(resp, &result)
		if err != nil {
			return &result, resp, fmt.Errorf("SalesReportsResource.GetSubscribersReports error: %v", err)
		}
		return &result, resp, fmt.Errorf(result.GetError())
	}
	return &result, resp, nil
}

//GetPreOrdersReports
func (srr *SalesReportsResource) GetPreOrdersReports(ctx context.Context, filter *PreOrdersReportsFilter) (*PreOrdersReportsResponse, *http.Response, error) {
	resp, err := srr.GetReports(ctx, filter)
	if err != nil {
		return nil, nil, fmt.Errorf("SalesReportsResource.GetPreOrdersReports error: %v", err)
	}
	result := PreOrdersReportsResponse{ResponseBody: &ResponseBody{}}
	result.status = resp.StatusCode
	if result.IsSuccess() {
		reports := []*PreOrdersReport{}
		err = srr.unmarshalResponse(resp, &reports)
		if err != nil {
			return &result, resp, fmt.Errorf("SalesReportsResource.GetPreOrdersReports error: %v", err)
		}
		result.Data = reports
	} else {
		err = srr.unmarshalResponse(resp, &result)
		if err != nil {
			return &result, resp, fmt.Errorf("SalesReportsResource.GetPreOrdersReports error: %v", err)
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
