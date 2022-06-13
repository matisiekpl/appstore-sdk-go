package appstore

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ConfigTestSuite struct {
	suite.Suite
}

func (suite *ConfigTestSuite) TestNewConfig() {
	result := NewConfig("foo", "bar", "baz", "tmp/key.p8")
	assert.Equal(suite.T(), AppStoreConnectAPIProductionUri, result.Uri)
	assert.Equal(suite.T(), "foo", result.IssuerId)
	assert.Equal(suite.T(), "bar", result.KeyId)
	assert.Equal(suite.T(), "baz", result.VendorNo)
	assert.Equal(suite.T(), "tmp/key.p8", result.PrivateKey)
	assert.Equal(suite.T(), "JWT", result.Token.Type)
	assert.Equal(suite.T(), "ES256", result.Token.Algo)
	assert.Equal(suite.T(), AppStoreConnectAPIAudience, result.Token.Audience)
	assert.Equal(suite.T(), AppStoreConnectAPITokenTtl, result.Token.Ttl)
}

func (suite *ConfigTestSuite) TestNewTokenConfig() {
	result := NewTokenConfig()
	assert.Equal(suite.T(), "JWT", result.Type)
	assert.Equal(suite.T(), "ES256", result.Algo)
	assert.Equal(suite.T(), AppStoreConnectAPIAudience, result.Audience)
	assert.Equal(suite.T(), AppStoreConnectAPITokenTtl, result.Ttl)
}

func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
