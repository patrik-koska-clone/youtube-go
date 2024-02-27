package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ApiKey string `yaml:"apiKey"`

	ChannelConfig struct {
		ChannelList []string `yaml:"channelList"`
	} `yaml:"channelConfig"`
}

func ReadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open config file\n%v", err)
	}

	var c Config

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal config\n%v", err)
	}

	return &c, nil
}
