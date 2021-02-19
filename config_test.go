package appstore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Config_NewConfig(t *testing.T) {
	result := NewConfig("foo", "bar", "baz", "tmp/key.p8")
	assert.Equal(t, AppStoreConnectAPIProductionUri, result.Uri)
	assert.Equal(t, "foo", result.IssuerId)
	assert.Equal(t, "bar", result.KeyId)
	assert.Equal(t, "baz", result.VendorNo)
	assert.Equal(t, "tmp/key.p8", result.PrivateKey)
	assert.Equal(t, "JWT", result.Token.Type)
	assert.Equal(t, "ES256", result.Token.Algo)
	assert.Equal(t, AppStoreConnectAPIAudience, result.Token.Audience)
	assert.Equal(t, AppStoreConnectAPITokenTtl, result.Token.Ttl)
}

func Test_Config_NewTokenConfig(t *testing.T) {
	result := NewTokenConfig()
	assert.Equal(t, "JWT", result.Type)
	assert.Equal(t, "ES256", result.Algo)
	assert.Equal(t, AppStoreConnectAPIAudience, result.Audience)
	assert.Equal(t, AppStoreConnectAPITokenTtl, result.Ttl)
}
