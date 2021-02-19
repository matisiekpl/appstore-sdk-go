package appstore

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

//AuthToken auth token structure
type AuthToken struct {
	Token     string
	ExpiresAt int64
}

//IsValid Check token is valid
func (t *AuthToken) IsValid() bool {
	return t.IsNotExpired() && t.Token != ""
}

//IsNotExpired Check token is not expired
func (t *AuthToken) IsNotExpired() bool {
	ts := time.Now().Unix()
	return t.ExpiresAt > ts
}

//TokenBuilder token builder
type TokenBuilder struct {
	cfg        *Config
	PrivateKey *PrivateKey
}

//NewTokenBuilder Create new TokenBuilder from config
func NewTokenBuilder(cfg *Config) *TokenBuilder {
	return &TokenBuilder{cfg: cfg, PrivateKey: &PrivateKey{}}
}

//BuildPayload Build JWT token payload
func (tb *TokenBuilder) BuildPayload() *jwt.StandardClaims {
	return &jwt.StandardClaims{
		Audience:  tb.cfg.Token.Audience,
		Issuer:    tb.cfg.IssuerId,
		ExpiresAt: time.Now().Unix() + int64(tb.cfg.Token.Ttl),
	}
}

//BuildJWTToken Build JWT token
func (tb *TokenBuilder) BuildJWTToken(payload *jwt.StandardClaims) *jwt.Token {
	return &jwt.Token{
		Header: map[string]interface{}{
			"typ": tb.cfg.Token.Type,
			"alg": tb.cfg.Token.Algo,
			"kid": tb.cfg.KeyId,
		},
		Claims: payload,
		Method: jwt.SigningMethodES256,
	}
}

//BuildAuthToken Build Auth token
func (tb *TokenBuilder) BuildAuthToken() (*AuthToken, error) {
	payload := tb.BuildPayload()
	jwtToken := tb.BuildJWTToken(payload)
	key, err := tb.PrivateKey.Load(tb.cfg.PrivateKey)
	if err != nil {
		return nil, err
	}
	secretStr, err := jwtToken.SignedString(key)
	if err != nil {
		return nil, err
	}
	token := &AuthToken{Token: secretStr, ExpiresAt: payload.ExpiresAt}
	return token, nil
}
