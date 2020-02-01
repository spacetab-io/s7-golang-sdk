package configuration

import (
	"crypto/tls"
)

type Config struct {
	Secret      secret
	Client      client
	Certificate tls.Certificate
}

type secret struct {
	Cert     string
	Key      string
	Login    string
	Password string
}

type client struct {
	URL        string `yaml:"url"`
	APIVersion string `yaml:"api_version"`
	Timeout    int64  `yaml:"timeout"`
}

func Initialization() (*Config, error) {

	conf := new(Config)



	certificate, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	conf.Certificate = certificate



	return conf, nil
}
