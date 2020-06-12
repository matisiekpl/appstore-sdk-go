package appstore

type Client struct {
	Cfg *Config
	TS  *TokenStorage
}

func NewClient(cfg *Config) *Client {
	cl := &Client{Cfg: cfg}
	return cl
}

type Token struct {
	Value     string
	ExpiresAt int64
}

type TokenStorage struct {
	Token *Token
}

type ResponseParser struct {
}
