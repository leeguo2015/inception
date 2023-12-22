package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Mysql  Mysql  `yaml:"mysql"`
	Logger Logger `yaml:"logger"`
	System System `yaml:"system"`
	Redis  Redis  `yaml:"redis"`
	Oss    Oss    `yaml:"oss"`
}

func New() *Config {
	return &Config{}
}
func (c *Config) Load(url string) error {

	yamlFile, err := os.ReadFile(url)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(yamlFile, c); err != nil {
		return err
	}
	return nil
}
