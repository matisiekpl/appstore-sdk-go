package appstore_sdk

const AppStoreConnectAPIProductionUri = "https://api.appstoreconnect.apple.com"

const AppStoreConnectAPIAudience = "appstoreconnect-v1"

type Config struct {
	Uri            string
	VendorNo       string
	IssuerId       string
	KeyId          string
	PrivateKeyPath string
	Token          *TokenConfig
}

type TokenConfig struct {
	Audience string
	Type     string
	Algo     string
	Ttl      int
}

func NewConfig(issuerId string, keyId string, vendorNo string, pkPath string) *Config {
	cfg := &Config{
		Uri:            AppStoreConnectAPIProductionUri,
		IssuerId:       issuerId,
		KeyId:          keyId,
		VendorNo:       vendorNo,
		PrivateKeyPath: pkPath,
		Token:          NewTokenConfig(),
	}
	return cfg
}

func NewTokenConfig() *TokenConfig {
	cfg := &TokenConfig{
		Type:     "JWT",
		Algo:     "ES256",
		Audience: AppStoreConnectAPIAudience,
		Ttl:      600,
	}
	return cfg
}
