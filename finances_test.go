package appstore

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func Test_Finances_FinancesReportsResource_GetReports_Success(t *testing.T) {
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

func Test_Finances_FinancesReportsResource_GetReports_WrongFilter(t *testing.T) {
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

func Test_Finances_FinancesReportsResource_GetFinancialReports_Success(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := buildStubConfig()
	transport := buildStubHttpTransport()
	ar := newResourceAbstract(transport, config)

	rsp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/finances/financial.tsv")
	rsp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", config.Uri+"/v1/financeReports", httpmock.ResponderFromResponse(rsp))

	resource := &FinancesReportsResource{ar}
	date, _ := time.Parse("2006-01-02", "2020-05-04")
	filter := &FinancesReportsFilter{ReportDate: date, RegionCode: "US", ReportType: FinancesReportTypeFinancial}
	ctx := context.Background()
	result, resp, err := resource.GetFinancialReports(ctx, filter)
	assert.NoError(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)

	assert.True(t, result.IsSuccess())
	assert.Empty(t, result.GetError())
	assert.Empty(t, result.Errors)
	assert.Equal(t, "2020-10-05", result.Data[0].StartDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, "2021-10-05", result.Data[0].EndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(t, "", result.Data[0].UPC)
	assert.Equal(t, "", result.Data[0].ISRCIsbn)
	assert.Equal(t, "foo.bar.baz", result.Data[0].VendorIdentifier)
	assert.Equal(t, 1, result.Data[0].Quantity.Value())
	assert.Equal(t, 3.1500000953674316, result.Data[0].PartnerShare.Value())
	assert.Equal(t, 3.1500000953674316, result.Data[0].ExtendedPartnerShare.Value())
	assert.Equal(t, "USD", result.Data[0].PartnerShareCurrency)
	assert.Equal(t, "S", result.Data[0].SaleOrReturn)
	assert.Equal(t, 1234567890, result.Data[0].AppleIdentifier.Value())
	assert.Equal(t, "", result.Data[0].ArtistShowDeveloperAuthor)
	assert.Equal(t, "foo.bar.baz", result.Data[0].Title)
	assert.Equal(t, "", result.Data[0].LabelStudioNetworkDeveloperPublisher)
	assert.Equal(t, "", result.Data[0].Grid)
	assert.Equal(t, "IAY", result.Data[0].ProductTypeIdentifier)
	assert.Equal(t, "", result.Data[0].ISANOtherIdentifier)
	assert.Equal(t, "US", result.Data[0].CountryOfSale)
	assert.Equal(t, "", result.Data[0].PreOrderFlag)
	assert.Equal(t, "", result.Data[0].PromoCode)
	assert.Equal(t, 4.489999771118164, result.Data[0].CustomerPrice.Value())
	assert.Equal(t, "USD", result.Data[0].CustomerCurrency)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(t, body)
}

func Test_Finances_FinancesReportsResource_GetFinancialReports_Error(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	config := buildStubConfig()
	transport := buildStubHttpTransport()
	ar := newResourceAbstract(transport, config)

	rsp := buildStubResponseFromFile(http.StatusBadRequest, "stubs/errors/invalid.parameter.json")
	rsp.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder("GET", config.Uri+"/v1/financeReports", httpmock.ResponderFromResponse(rsp))

	resource := &FinancesReportsResource{ar}
	date, _ := time.Parse("2006-01-02", "2020-05-04")
	filter := &FinancesReportsFilter{ReportDate: date, RegionCode: "US", ReportType: FinancesReportTypeFinancial}
	ctx := context.Background()
	result, resp, err := resource.GetFinancialReports(ctx, filter)
	assert.Error(t, err)
	assert.NotEmpty(t, resp)
	assert.NotEmpty(t, result)
	assert.Equal(t, "The version parameter you have specified is invalid. The latest version for this report is 1_0.", err.Error())

	assert.False(t, result.IsSuccess())
	assert.Equal(t, "The version parameter you have specified is invalid. The latest version for this report is 1_0.", result.GetError())
	assert.Len(t, result.Errors, 1)
	assert.Empty(t, result.Data)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(t, body)
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
