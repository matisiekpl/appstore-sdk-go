package appstore_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_PrivateKey_LoadFromFile(t *testing.T) {
	pk := &PrivateKey{}
	data, _ := pk.LoadFromFile(StubAuthKeyPath)
	assert.NotEmpty(t, data)
}

func Test_PrivateKey_LoadFromContent(t *testing.T) {
	pk := &PrivateKey{}
	data, _ := pk.LoadFromFile(StubAuthKeyPath)
	data, _ = pk.LoadFromContent(string(data))
	assert.NotEmpty(t, data)
}

func Test_PrivateKey_LoadSuccess(t *testing.T) {
	pk := &PrivateKey{}
	key, _ := pk.Load(StubAuthKeyPath)
	assert.NotEmpty(t, key)
}

func Test_PrivateKey_LoadDataAsFile(t *testing.T) {
	pk := &PrivateKey{}
	data, _ := pk.LoadData(StubAuthKeyPath)
	assert.NotEmpty(t, data)
}

func Test_PrivateKey_LoadDataAsContent(t *testing.T) {
	pk := &PrivateKey{}
	data, _ := pk.LoadFromFile(StubAuthKeyPath)
	data, _ = pk.LoadData(string(data))
	assert.NotEmpty(t, data)
}

func Test_PrivateKey_ParseP8Success(t *testing.T) {
	pk := &PrivateKey{}
	data, _ := pk.LoadFromFile(StubAuthKeyPath)
	key, _ := pk.ParseP8(data)
	assert.NotEmpty(t, key)
}

func Test_PrivateKey_DecodePemSuccess(t *testing.T) {
	pk := &PrivateKey{}
	data, _ := pk.LoadFromFile(StubAuthKeyPath)
	pem, _ := pk.DecodePem(data)
	assert.NotEmpty(t, pem)
}
