package appstore

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func Test_Finances_FinancesReportsResource_Success(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := buildStubConfig()
	transport := buildStubHttpTransport()

	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/finances/financial.tsv")
	resp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", config.Uri+"/v1/financeReports", httpmock.ResponderFromResponse(resp))

	resource := &FinancesReportsResource{ResourceAbstract: newResourceAbstract(transport, config)}
	date, _ := time.Parse("2006-01-02", "2020-05-04")
	filter := &FinancesReportsFilter{ReportDate: date, RegionCode: "US", ReportType: FinancesReportTypeFinancial}
	ctx := context.Background()
	rsp, err := resource.GetReports(ctx, filter)
	assert.NoError(t, err)
	assert.NotEmpty(t, rsp)
}

func Test_Finances_FinancesReportsResource_WrongFilter(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := buildStubConfig()
	transport := buildStubHttpTransport()

	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/finances/financial.tsv")
	resp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", config.Uri+"/v1/financeReports", httpmock.ResponderFromResponse(resp))

	resource := &FinancesReportsResource{ResourceAbstract: newResourceAbstract(transport, config)}
	date, _ := time.Parse("2006-01-02", "2020-05-04")
	filter := &FinancesReportsFilter{ReportDate: date, RegionCode: "US"}
	ctx := context.Background()
	rsp, err := resource.GetReports(ctx, filter)
	assert.Error(t, err)
	assert.Empty(t, rsp)
	assert.Equal(t, "FinancesReportsResource.GetReports invalid filter: FinancesReportsFilter.IsValid: ReportType is required", err.Error())
}

func Test_Finances_FinancesReportsResource_BuildQueryParams(t *testing.T) {
	config := buildStubConfig()
	token := buildStubAuthToken()
	transport := NewHttpTransport(config, token, nil)
	resource := &FinancesReportsResource{ResourceAbstract: newResourceAbstract(transport, config)}
	date, _ := time.Parse("2006-01-02", "2020-05-04")
	filter := &FinancesReportsFilter{ReportDate: date, RegionCode: "US", ReportType: FinancesReportTypeFinancial}
	result := resource.buildQueryParams(filter)
	qs := make(map[string]interface{})
	qs["filter[reportDate]"] = "2020-05"
	qs["filter[reportType]"] = string(FinancesReportTypeFinancial)
	qs["filter[regionCode]"] = "US"
	qs["filter[vendorNumber]"] = config.VendorNo
	assert.Equal(t, qs, result)
}
