package appstore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Client_NewClientFromConfig(t *testing.T) {
	config := buildStubConfig()
	result := NewClientFromConfig(config, nil)
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.Cfg)
	assert.NotEmpty(t, result.http)
	assert.NotEmpty(t, result.auth)
	assert.Empty(t, result.transport)
}

func Test_Client_InitSuccess(t *testing.T) {
	config := buildStubConfig()
	result := NewClientFromConfig(config, nil)
	_ = result.Init()
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.Cfg)
	assert.NotEmpty(t, result.http)
	assert.NotEmpty(t, result.auth)
	assert.NotEmpty(t, result.transport)
}

func Test_Client_SalesReports(t *testing.T) {
	config := buildStubConfig()
	client := NewClientFromConfig(config, nil)
	_ = client.Init()
	result := client.SalesReports()
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.config)
	assert.NotEmpty(t, result.transport)
}

func Test_Client_FinancialReports(t *testing.T) {
	config := buildStubConfig()
	client := NewClientFromConfig(config, nil)
	_ = client.Init()
	result := client.FinancesReports()
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.config)
	assert.NotEmpty(t, result.transport)
}

func Test_Client_InitError(t *testing.T) {
	config := buildStubConfig()
	config.PrivateKey = "stubs/auth/keys/fail.p8"
	result := NewClientFromConfig(config, nil)
	err := result.Init()
	assert.Error(t, err)
	assert.Equal(t, "client.init error: PrivateKey.DecodePem: AuthKey must be a valid .p8 PEM file", err.Error())
}
