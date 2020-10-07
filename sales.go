package appstore_sdk

type SalesReportsResource struct {
	*ResourceAbstract
}

func (srr *SalesReportsResource) GetReport(filter *SalesReportsFilter) (*Response, error) {
	return srr.get("v1/salesReports", filter.ToQueryParamsMap())
}
