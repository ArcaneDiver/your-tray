package trays

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Tray Tray `yaml:"tray,flow"`
}

func Parse(path string) (*Config, error) {
	config := Config{}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	if err := yaml.UnmarshalStrict(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}