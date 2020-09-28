package appstore_sdk

type Client struct {
	Cfg *Config
	Ts  *TokenStorage
	Pk  *PrivateKey
}

func NewClient(cfg *Config) *Client {
	cl := &Client{Cfg: cfg, Ts: &TokenStorage{}, Pk: &PrivateKey{}}
	return cl
}

type Token struct {
	Value     string
	ExpiresAt int64
}

type TokenStorage struct {
	Token *Token
}
