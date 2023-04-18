package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Info struct {
		Author  string
		Version string
	}
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
	Lavalink struct {
		Name     string `yaml:"name"`
		Address  string `yaml:"address"`
		Password string `yaml:"password"`
		Https    bool   `yaml:"https"`
	}
	Spotify struct {
		Client, Secret string
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
