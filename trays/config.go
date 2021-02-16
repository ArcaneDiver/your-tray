package trays

import (
	"github.com/ArcaneDiver/your-tray/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Tray 		Tray `yaml:"tray,flow"`
	UpdateRate	int  `yaml:"updateRate"`
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

	if config.UpdateRate <= 0 {
		log.Log.Warn("updateRate cannot be lower than 1, setting to 1")
		config.UpdateRate = 1
	}

	log.Log.WithField("config", &config).Debug("parsed config")
	return &config, nil
}