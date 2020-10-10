package appstore_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Client_NewClientFromConfig(t *testing.T) {
	config := BuildStubConfig()
	token := BuildStubAuthToken()
	result := NewClientFromConfig(config, token, nil)
	assert.NotEmpty(t, result)
}
