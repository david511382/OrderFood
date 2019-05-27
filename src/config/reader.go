package config

import (
	"orderfood/src/util"

	"gopkg.in/yaml.v2"
)

// ReadConfig read config from filepath
func ReadConfig(filename string) error {
	cfgBytes, err := util.ReadFile(filename)
	if err != nil {
		return err
	}

	cfg = &Config{}
	err = yaml.Unmarshal(cfgBytes, cfg)

	return err
}
