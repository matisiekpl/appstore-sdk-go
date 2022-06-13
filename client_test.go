package appstore

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ClientTestSuite struct {
	suite.Suite
	cfg      *Config
	testable *Client
}

func (suite *ClientTestSuite) SetupTest() {
	suite.cfg = buildStubConfig()
	suite.testable = NewClientFromConfig(suite.cfg, nil)
}

func (suite *ClientTestSuite) TestNewClientFromConfig() {
	assert.NotEmpty(suite.T(), suite.testable.Cfg)
	assert.NotEmpty(suite.T(), suite.testable.http)
	assert.NotEmpty(suite.T(), suite.testable.auth)
	assert.Empty(suite.T(), suite.testable.transport)
}

func (suite *ClientTestSuite) TestInitSuccess() {
	err := suite.testable.Init()
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), suite.testable.Cfg)
	assert.NotEmpty(suite.T(), suite.testable.http)
	assert.NotEmpty(suite.T(), suite.testable.auth)
	assert.NotEmpty(suite.T(), suite.testable.transport)
}

func (suite *ClientTestSuite) TestInitError() {
	suite.testable.Cfg.PrivateKey = "stubs/auth/keys/fail.p8"
	err := suite.testable.Init()
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "client.init error: PrivateKey.DecodePem: AuthKey must be a valid .p8 PEM file", err.Error())
}

func (suite *ClientTestSuite) TestSalesReports() {
	_ = suite.testable.Init()
	result := suite.testable.SalesReports()
	assert.NotEmpty(suite.T(), result)
	assert.NotEmpty(suite.T(), result.config)
	assert.NotEmpty(suite.T(), result.transport)
}

func (suite *ClientTestSuite) TestFinancialReports() {
	_ = suite.testable.Init()
	result := suite.testable.FinancesReports()
	assert.NotEmpty(suite.T(), result)
	assert.NotEmpty(suite.T(), result.config)
	assert.NotEmpty(suite.T(), result.transport)
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
