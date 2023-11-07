package config

var Conf Config

type Config struct {
	Server ServerConfig `yaml:"ServerConfig"`
}

type ServerConfig struct {
	Port    string `yaml:"Port"`
	UrlPath string `yaml:"UrlPath"`
}
