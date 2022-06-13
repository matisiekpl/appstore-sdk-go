package appstore

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type PrivateKeyTestSuite struct {
	suite.Suite
	testable *PrivateKey
}

func (suite *PrivateKeyTestSuite) SetupTest() {
	suite.testable = &PrivateKey{}
}

func (suite *PrivateKeyTestSuite) TestLoadFromFile() {
	data, err := suite.testable.LoadFromFile(StubAuthKeyPath)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), data)
}

func (suite *PrivateKeyTestSuite) TestLoadFromContent() {
	data, _ := suite.testable.LoadFromFile(StubAuthKeyPath)
	data, err := suite.testable.LoadFromContent(string(data))
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), data)
}

func (suite *PrivateKeyTestSuite) TestLoadSuccess() {
	key, err := suite.testable.Load(StubAuthKeyPath)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), key)
}

func (suite *PrivateKeyTestSuite) TestLoadDataAsFile() {
	data, err := suite.testable.LoadData(StubAuthKeyPath)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), data)
}

func (suite *PrivateKeyTestSuite) TestLoadDataAsContent() {
	dataBt, err := suite.testable.LoadData("foo")
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), dataBt)
}

func (suite *PrivateKeyTestSuite) TestParseP8Success() {
	data, _ := suite.testable.LoadFromFile(StubAuthKeyPath)
	key, err := suite.testable.ParseP8(data)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), key)
}

func (suite *PrivateKeyTestSuite) TestDecodePemSuccess() {
	data, _ := suite.testable.LoadFromFile(StubAuthKeyPath)
	pem, err := suite.testable.DecodePem(data)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), pem)
}

func TestPrivateKeyTestSuite(t *testing.T) {
	suite.Run(t, new(PrivateKeyTestSuite))
}
