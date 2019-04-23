package jwt

var config Config

type Config struct {
	Secret    string
	secretKey []byte
}

func Setup(cfg Config) {
	cfg.secretKey = []byte(cfg.Secret)
	config = cfg
}
