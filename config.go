package appstore

import "time"

//AppStoreConnectAPIProductionUri const
const AppStoreConnectAPIProductionUri = "https://api.appstoreconnect.apple.com"

//AppStoreConnectAPIAudience const
const AppStoreConnectAPIAudience = "appstoreconnect-v1"

//AppStoreConnectAPITokenTtl const
const AppStoreConnectAPITokenTtl = 600

//AppStoreConnectAPIHttpMaxIdleConnection const
const AppStoreConnectAPIHttpMaxIdleConnection = 10

//AppStoreConnectAPIHttpIdleConnectionTimeout const
const AppStoreConnectAPIHttpIdleConnectionTimeout = 30 * time.Second

//Config structure
type Config struct {
	Uri        string
	VendorNo   string
	IssuerId   string
	KeyId      string
	PrivateKey string
	Token      *TokenConfig
}

//TokenConfig token config structure
type TokenConfig struct {
	Audience string
	Type     string
	Algo     string
	Ttl      int
}

//NewConfig Create new config from credentials
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

//NewTokenConfig Create new token config
func NewTokenConfig() *TokenConfig {
	cfg := &TokenConfig{
		Type:     "JWT",
		Algo:     "ES256",
		Audience: AppStoreConnectAPIAudience,
		Ttl:      AppStoreConnectAPITokenTtl,
	}
	return cfg
}
