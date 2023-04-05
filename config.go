package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Discord struct {
		Token string `yaml:"token"`
		Guild string `yaml:"guild"`
		Rmcmd bool   `yaml:"rmcmd"`
	}
	Database struct {
		Url        string   `yaml:"url"`
		Database   string   `yaml:"database"`
		Collection []string `yaml:"collection"`
	}
}

func Config_init() (*Config, error) {
	buf, err := ioutil.ReadFile("config.yml")
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", "config.yml", err)
	}

	return c, err
}
