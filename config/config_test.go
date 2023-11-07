package config

import (
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestLoadConfigRaw(t *testing.T) {
	conf := &Config{}
	content, err := os.ReadFile("config.yml")
	if err != nil {
		t.Error(err)
	}
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		t.Error(err)
	}
	if conf.Server.Port == "" {
		t.Error("Port == \"\"")
	}
	t.Log(conf)
}

func TestLoadConfig(t *testing.T) {
	c := &Config{}
	MustLoadConfig("config.yml", &c)
	if c.Server.Port == "" {
		t.Error("Port == \"\"")
	}
	t.Log(c)
}
