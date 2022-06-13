package appstore

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type AuthTokenTestSuite struct {
	suite.Suite
	testable *AuthToken
}

func (suite *AuthTokenTestSuite) SetupTest() {
	suite.testable = &AuthToken{}
}

func (suite *AuthTokenTestSuite) TestIsNotExpiredSuccess() {
	suite.testable.ExpiresAt = time.Now().Unix() + 1000
	assert.True(suite.T(), suite.testable.IsNotExpired())
}

func (suite *AuthTokenTestSuite) TestIsNotExpiredFail() {
	suite.testable.ExpiresAt = time.Now().Unix() - 1000
	assert.False(suite.T(), suite.testable.IsNotExpired())
}

func (suite *AuthTokenTestSuite) TestIsValidSuccess() {
	suite.testable.ExpiresAt = time.Now().Unix() + 1000
	suite.testable.Token = "qwerty"
	assert.True(suite.T(), suite.testable.IsValid())
}

func (suite *AuthTokenTestSuite) TestIsValidFailExpired() {
	suite.testable.ExpiresAt = time.Now().Unix() - 1000
	suite.testable.Token = "qwerty"
	assert.False(suite.T(), suite.testable.IsValid())
}

func (suite *AuthTokenTestSuite) TestIsValidFailEmpty() {
	suite.testable.ExpiresAt = time.Now().Unix() + 1000
	assert.False(suite.T(), suite.testable.IsValid())
}

func TestAuthTokenTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTokenTestSuite))
}

type AuthTokenBuilderTestSuite struct {
	suite.Suite
	cfg      *Config
	testable *TokenBuilder
}

func (suite *AuthTokenBuilderTestSuite) SetupTest() {
	suite.cfg = buildStubConfig()
	suite.testable = NewTokenBuilder(suite.cfg)
}

func (suite *AuthTokenBuilderTestSuite) TestBuildPayload() {
	payload := suite.testable.BuildPayload()
	assert.NotEmpty(suite.T(), payload)
	assert.NotEmpty(suite.T(), payload.ExpiresAt)
	assert.Equal(suite.T(), suite.cfg.Token.Audience, payload.Audience)
	assert.Equal(suite.T(), suite.cfg.IssuerId, payload.Issuer)
}

func (suite *AuthTokenBuilderTestSuite) TestBuildJWTToken() {
	payload := suite.testable.BuildPayload()
	jwtToken := suite.testable.BuildJWTToken(payload)
	assert.NotEmpty(suite.T(), jwtToken)
	assert.Equal(suite.T(), jwtToken.Claims, payload)
	assert.Equal(suite.T(), jwtToken.Method, jwt.SigningMethodES256)
	assert.Equal(suite.T(), jwtToken.Header["typ"], suite.cfg.Token.Type)
	assert.Equal(suite.T(), jwtToken.Header["alg"], suite.cfg.Token.Algo)
	assert.Equal(suite.T(), jwtToken.Header["kid"], suite.cfg.KeyId)
}

func (suite *AuthTokenBuilderTestSuite) TestBuildAuthTokenSuccess() {
	token, err := suite.testable.BuildAuthToken()
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), token)
	assert.NotEmpty(suite.T(), token.Token)
	assert.NotEmpty(suite.T(), token.ExpiresAt)
	assert.True(suite.T(), token.IsValid())
	assert.True(suite.T(), token.IsNotExpired())
}

func (suite *AuthTokenBuilderTestSuite) TestBuildAuthTokenError() {
	suite.cfg.PrivateKey = "stubs/auth/keys/fail.p8"
	token, err := suite.testable.BuildAuthToken()
	assert.Nil(suite.T(), token)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "PrivateKey.DecodePem: AuthKey must be a valid .p8 PEM file", err.Error())
}

func TestAuthTokenBuilderTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTokenBuilderTestSuite))
}
