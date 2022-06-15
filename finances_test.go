package appstore

import (
	"context"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

type FinancesReportsResourceTestSuite struct {
	suite.Suite
	cfg      *Config
	ctx      context.Context
	testable *FinancesReportsResource
}

func (suite *FinancesReportsResourceTestSuite) SetupTest() {
	suite.cfg = buildStubConfig()
	suite.ctx = context.Background()
	suite.testable = buildStubFinancesReportsResource()
	httpmock.Activate()
}

func (suite *FinancesReportsResourceTestSuite) TearDownTest() {
	httpmock.DeactivateAndReset()
}

func (suite *FinancesReportsResourceTestSuite) TestGetReportsSuccess() {
	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/finances/financial.tsv")
	resp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/v1/financeReports", httpmock.ResponderFromResponse(resp))

	date, _ := time.Parse("2006-01-02", "2020-05-04")
	filter := &FinancesReportsFilter{ReportDate: date, RegionCode: "US", ReportType: FinancesReportTypeFinancial}
	rsp, err := suite.testable.GetReports(suite.ctx, filter)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), rsp)
}

func (suite *FinancesReportsResourceTestSuite) TestGetReportsWrongFilter() {
	resp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/finances/financial.tsv")
	resp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/v1/financeReports", httpmock.ResponderFromResponse(resp))

	date, _ := time.Parse("2006-01-02", "2020-05-04")
	filter := &FinancesReportsFilter{ReportDate: date, RegionCode: "US"}
	rsp, err := suite.testable.GetReports(suite.ctx, filter)
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), rsp)
	assert.Equal(suite.T(), "FinancesReportsResource.GetReports invalid filter: FinancesReportsFilter.IsValid: ReportType is required", err.Error())
}

func (suite *FinancesReportsResourceTestSuite) TestGetFinancialReportsSuccess() {
	rsp := buildStubResponseFromGzip(http.StatusOK, "stubs/reports/finances/financial.tsv")
	rsp.Header.Set("Content-Type", ResponseContentTypeGzip)
	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/v1/financeReports", httpmock.ResponderFromResponse(rsp))

	date, _ := time.Parse("2006-01-02", "2020-05-04")
	filter := &FinancesReportsFilter{ReportDate: date, RegionCode: "US", ReportType: FinancesReportTypeFinancial}
	result, resp, err := suite.testable.GetFinancialReports(suite.ctx, filter)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)

	assert.True(suite.T(), result.IsSuccess())
	assert.Empty(suite.T(), result.GetError())
	assert.Empty(suite.T(), result.Errors)
	assert.Equal(suite.T(), "2020-10-05", result.Data[0].StartDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), "2021-10-05", result.Data[0].EndDate.Value().Format(CustomDateFormatDefault))
	assert.Equal(suite.T(), "", result.Data[0].UPC)
	assert.Equal(suite.T(), "", result.Data[0].ISRCIsbn)
	assert.Equal(suite.T(), "foo.bar.baz", result.Data[0].VendorIdentifier)
	assert.Equal(suite.T(), 1, result.Data[0].Quantity.Value())
	assert.Equal(suite.T(), 3.1500000953674316, result.Data[0].PartnerShare.Value())
	assert.Equal(suite.T(), 3.1500000953674316, result.Data[0].ExtendedPartnerShare.Value())
	assert.Equal(suite.T(), "USD", result.Data[0].PartnerShareCurrency)
	assert.Equal(suite.T(), "S", result.Data[0].SaleOrReturn)
	assert.Equal(suite.T(), 1234567890, result.Data[0].AppleIdentifier.Value())
	assert.Equal(suite.T(), "", result.Data[0].ArtistShowDeveloperAuthor)
	assert.Equal(suite.T(), "foo.bar.baz", result.Data[0].Title)
	assert.Equal(suite.T(), "", result.Data[0].LabelStudioNetworkDeveloperPublisher)
	assert.Equal(suite.T(), "", result.Data[0].Grid)
	assert.Equal(suite.T(), "IAY", result.Data[0].ProductTypeIdentifier)
	assert.Equal(suite.T(), "", result.Data[0].ISANOtherIdentifier)
	assert.Equal(suite.T(), "US", result.Data[0].CountryOfSale)
	assert.Equal(suite.T(), "", result.Data[0].PreOrderFlag)
	assert.Equal(suite.T(), "", result.Data[0].PromoCode)
	assert.Equal(suite.T(), 4.489999771118164, result.Data[0].CustomerPrice.Value())
	assert.Equal(suite.T(), "USD", result.Data[0].CustomerCurrency)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
}

func (suite *FinancesReportsResourceTestSuite) TestGetFinancialReportsError() {
	rsp := buildStubResponseFromFile(http.StatusBadRequest, "stubs/errors/invalid.parameter.json")
	rsp.Header.Set("Content-Type", ResponseContentTypeJson)
	httpmock.RegisterResponder("GET", suite.cfg.Uri+"/v1/financeReports", httpmock.ResponderFromResponse(rsp))

	date, _ := time.Parse("2006-01-02", "2020-05-04")
	filter := &FinancesReportsFilter{ReportDate: date, RegionCode: "US", ReportType: FinancesReportTypeFinancial}
	result, resp, err := suite.testable.GetFinancialReports(suite.ctx, filter)
	assert.Error(suite.T(), err)
	assert.NotEmpty(suite.T(), resp)
	assert.NotEmpty(suite.T(), result)
	assert.Equal(suite.T(), "The version parameter you have specified is invalid. The latest version for this report is 1_0.", err.Error())

	assert.False(suite.T(), result.IsSuccess())
	assert.Equal(suite.T(), "The version parameter you have specified is invalid. The latest version for this report is 1_0.", result.GetError())
	assert.Len(suite.T(), result.Errors, 1)
	assert.Empty(suite.T(), result.Data)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.NotEmpty(suite.T(), body)
}

func (suite *FinancesReportsResourceTestSuite) TestGetFinancialReportsErrorWrongFilter() {
	date, _ := time.Parse("2006-01-02", "2020-05-04")
	filter := &FinancesReportsFilter{ReportDate: date, RegionCode: "US"}

	result, resp, err := suite.testable.GetFinancialReports(suite.ctx, filter)
	assert.Error(suite.T(), err)
	assert.Empty(suite.T(), resp)
	assert.Empty(suite.T(), result)
	assert.Equal(suite.T(), "FinancialReportsResponse.GetFinancialReports error: FinancesReportsResource.GetReports invalid filter: FinancesReportsFilter.IsValid: ReportType is required", err.Error())
}

func (suite *FinancesReportsResourceTestSuite) TestBuildQueryParams() {
	date, _ := time.Parse("2006-01-02", "2020-05-04")
	filter := &FinancesReportsFilter{ReportDate: date, RegionCode: "US", ReportType: FinancesReportTypeFinancial}
	result := suite.testable.buildQueryParams(filter)
	qs := make(map[string]interface{})
	qs["filter[reportDate]"] = "2020-05"
	qs["filter[reportType]"] = string(FinancesReportTypeFinancial)
	qs["filter[regionCode]"] = "US"
	qs["filter[vendorNumber]"] = suite.cfg.VendorNo
	assert.Equal(suite.T(), qs, result)
}

func TestFinancesReportsResourceTestSuite(t *testing.T) {
	suite.Run(t, new(FinancesReportsResourceTestSuite))
}
