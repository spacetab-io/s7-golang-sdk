package configuration

import (
	"crypto/tls"
)

type Config struct {
	Secret      Secret
	Client      Client
	Certificate tls.Certificate
}

type Secret struct {
	Cert     string
	Key      string
	Login    string
	Password string
}

type Client struct {
	URL        string
	APIVersion string
	Timeout    int64
}

func (c *Config) Initialization(config Config) error {

	certificate, err := tls.LoadX509KeyPair(config.Secret.Cert, config.Secret.Key)
	if err != nil {
		return err
	}

	c.Certificate = certificate

	return nil
}
