package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func MustLoadConfig(filePath string, conf any) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("[FATAl] func=MustLoadConfig, step=os.ReadFile, filePath=%s\n", filePath)
	}
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		log.Fatalf("[FATAl] func=MustLoadConfig, step=yaml.Unmarshal, filePath=%s, content=%v\n", filePath, conf)
	}
}
