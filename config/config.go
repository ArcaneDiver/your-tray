package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Trays 	struct {
		Name	string `yaml:"name"`
		Tooltip	string `yaml:"tooltip"`
		Items	[] struct {
			Text	string `yaml:"text"`
			Tooltip	string `yaml:"tooltip"`
			Command string `yaml:"command"`
		} `yaml:"items,flow"`
	} `yaml:"trays"`
	Icon	string `yaml:"icon"`
}

func Parse(path string) (*Config, error) {
	config := Config{}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}