package appstore

type Config struct {
	VendorNo       string
	IssuerId       string
	KeyId          string
	PrivateKey     string
	PrivateKeyPath string
	TokenTtl       int64
}

func NewConfig() *Config {
	cfg := &Config{}
	cfg.TokenTtl = 600
	return cfg
}

//func (cfg *Config) LoadPrivateKey() error {
//	if cfg.PrivateKey == "" {
//		privateKey, err := ioutil.ReadFile(cfg.PrivateKeyPath)
//	}
//	privateKey, err := ioutil.ReadFile(cfg.PrivateKeyPath)
//	return string(privateKey), err
//}
