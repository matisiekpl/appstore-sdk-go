package appstore_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Resource_NewResourceAbstract(t *testing.T) {
	config := BuildStubConfig()
	token := BuildStubAuthToken()
	transport := NewHttpTransport(config, token, nil)
	result := newResourceAbstract(transport, config)
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.config)
	assert.NotEmpty(t, result.transport)
}
