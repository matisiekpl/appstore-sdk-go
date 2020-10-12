package appstore_sdk

import "time"

const AppStoreConnectAPIProductionUri = "https://api.appstoreconnect.apple.com"
const AppStoreConnectAPIAudience = "appstoreconnect-v1"
const AppStoreConnectAPITokenTtl = 600
const AppStoreConnectAPIHttpMaxIdleConnection = 10
const AppStoreConnectAPIHttpIdleConnectionTimeout = 30 * time.Second

//Common config
type Config struct {
	Uri        string
	VendorNo   string
	IssuerId   string
	KeyId      string
	PrivateKey string
	Token      *TokenConfig
}

//Token config
type TokenConfig struct {
	Audience string
	Type     string
	Algo     string
	Ttl      int
}

//Create new config from credentials
func NewConfig(issuerId string, keyId string, vendorNo string, pkPathOrContent string) *Config {
	cfg := &Config{
		Uri:        AppStoreConnectAPIProductionUri,
		IssuerId:   issuerId,
		KeyId:      keyId,
		VendorNo:   vendorNo,
		PrivateKey: pkPathOrContent,
		Token:      NewTokenConfig(),
	}
	return cfg
}

//Create new token config
func NewTokenConfig() *TokenConfig {
	cfg := &TokenConfig{
		Type:     "JWT",
		Algo:     "ES256",
		Audience: AppStoreConnectAPIAudience,
		Ttl:      AppStoreConnectAPITokenTtl,
	}
	return cfg
}
