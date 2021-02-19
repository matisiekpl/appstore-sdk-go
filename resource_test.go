package appstore

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
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

func Test_Resources_Resource_Get(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	cfg := buildStubConfig()
	transport := buildStubHttpTransport()
	resource := newResourceAbstract(transport, cfg)

	body, _ := loadStubResponseData("stubs/reports/sales/sales.tsv")

	httpmock.RegisterResponder("GET", cfg.Uri+"/foo", httpmock.NewBytesResponder(http.StatusOK, body))

	resp, _ := resource.get("foo", nil)
	assert.NotEmpty(t, resp)
}
