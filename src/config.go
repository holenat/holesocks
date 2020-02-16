package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/pelletier/go-toml"
)

type Config struct {
	Addr  string            `toml:"addr"`
	Users map[string]string `toml:"users"`
}

func ParseConfig(path string) (*Config, error) {
	cnt, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = toml.Unmarshal(cnt, &cfg)
	return &cfg, err
}

func (c *Config) String() string {
	b, _ := json.MarshalIndent(c, "", "   ")
	return string(b)
}
