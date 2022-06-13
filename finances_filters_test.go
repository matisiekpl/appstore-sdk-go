package appstore

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type FinancesReportsFilterTestSuite struct {
	suite.Suite
}

func (suite *FinancesReportsFilterTestSuite) TestIsValid() {
	date, _ := time.Parse("2006-01-02", "2020-04-17")
	filter := NewFinancesReportsFilter()
	filter.SetReportDate(date).SetRegionCode("US")
	assert.NoError(suite.T(), filter.IsValid())
	assert.Equal(suite.T(), FinancesReportTypeFinancial, filter.ReportType)
}

func (suite *FinancesReportsFilterTestSuite) TestIsInvalidEmptyReportType() {
	date, _ := time.Parse("2006-01-02", "2020-04-17")
	filter := &FinancesReportsFilter{ReportDate: date, RegionCode: "US"}
	err := filter.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "FinancesReportsFilter.IsValid: ReportType is required", err.Error())
}

func (suite *FinancesReportsFilterTestSuite) TestIsInvalidEmptyRegionCode() {
	date, _ := time.Parse("2006-01-02", "2020-04-17")
	filter := &FinancesReportsFilter{ReportDate: date}
	filter.TypeFinancial()
	err := filter.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "FinancesReportsFilter.IsValid: RegionCode is required", err.Error())
}

func (suite *FinancesReportsFilterTestSuite) TestIsInvalidEmptyReportDate() {
	filter := &FinancesReportsFilter{RegionCode: "US"}
	filter.TypeFinanceDetail()
	err := filter.IsValid()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "FinancesReportsFilter.IsValid: ReportDate is required", err.Error())
}

func (suite *FinancesReportsFilterTestSuite) TestToQueryParamsMap() {
	date, _ := time.Parse("2006-01-02", "2020-05-04")
	filter := &FinancesReportsFilter{ReportDate: date, RegionCode: "US", ReportType: FinancesReportTypeFinancial}
	qs := make(map[string]interface{})
	qs["filter[reportDate]"] = "2020-05"
	qs["filter[reportType]"] = string(FinancesReportTypeFinancial)
	qs["filter[regionCode]"] = "US"
	assert.Equal(suite.T(), qs, filter.toQueryParamsMap())
}

func TestFinancesReportsFilterTestSuite(t *testing.T) {
	suite.Run(t, new(FinancesReportsFilterTestSuite))
}
