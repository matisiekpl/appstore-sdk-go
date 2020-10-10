package appstore_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Filesystem_fileExistsSuccess(t *testing.T) {
	exists := fileExists(StubAuthKeyPath)
	assert.True(t, exists)
}

func Test_Filesystem_readFileSuccess(t *testing.T) {
	exists, _ := readFile(StubAuthKeyPath)
	assert.NotEmpty(t, exists)
}
