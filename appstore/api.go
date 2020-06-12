package appstore

type Client struct {
	Cfg *Config
	Ts  *TokenStorage
}

func NewClient(cfg *Config) *Client {
	cl := &Client{Cfg: cfg, Ts: &TokenStorage{}}
	return cl
}

type Token struct {
	Value     string
	ExpiresAt int64
}

type TokenStorage struct {
	Token *Token
}
