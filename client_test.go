package appstore_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Client_NewClientFromConfig(t *testing.T) {
	config := BuildStubConfig()
	result := NewClientFromConfig(config, nil)
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.Cfg)
	assert.NotEmpty(t, result.http)
	assert.NotEmpty(t, result.auth)
	assert.Empty(t, result.transport)
}

func Test_Client_Init(t *testing.T) {
	config := BuildStubConfig()
	result := NewClientFromConfig(config, nil)
	_ = result.Init()
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.Cfg)
	assert.NotEmpty(t, result.http)
	assert.NotEmpty(t, result.auth)
	assert.NotEmpty(t, result.transport)
}
