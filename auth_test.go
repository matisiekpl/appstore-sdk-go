package appstore_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_Auth_TokenIsNotExpiredSuccess(t *testing.T) {
	expired := time.Now().Unix() + 1000
	token := AuthToken{ExpiresAt: expired}
	assert.True(t, token.IsNotExpired())
}

func Test_Auth_TokenIsNotExpiredFail(t *testing.T) {
	expired := time.Now().Unix() - 1000
	token := AuthToken{ExpiresAt: expired}
	assert.False(t, token.IsNotExpired())
}

func Test_Auth_TokenIsValidSuccess(t *testing.T) {
	expired := time.Now().Unix() + 1000
	token := AuthToken{ExpiresAt: expired, Token: "qwerty"}
	assert.True(t, token.IsValid())
}

func Test_Auth_TokenIsValidFailExpired(t *testing.T) {
	expired := time.Now().Unix() - 1000
	token := AuthToken{ExpiresAt: expired, Token: "qwerty"}
	assert.False(t, token.IsValid())
}

func Test_Auth_TokenIsValidFailEmpty(t *testing.T) {
	expired := time.Now().Unix() + 1000
	token := AuthToken{ExpiresAt: expired}
	assert.False(t, token.IsValid())
}
