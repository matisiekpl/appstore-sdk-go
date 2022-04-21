package appstore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Resource_NewResourceAbstract(t *testing.T) {
	config := buildStubConfig()
	token := buildStubAuthToken()
	transport := NewHttpTransport(config, token, nil)
	result := newResourceAbstract(transport, config)
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.config)
	assert.NotEmpty(t, result.transport)
}
