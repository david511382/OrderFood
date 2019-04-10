package config

import (
	"fmt"
	"io/ioutil"
	"orderfood/src/util"

	"gopkg.in/yaml.v2"
)

// ReadConfig read config from filepath
func ReadConfig(filename string) error {
	filepath, err := util.GetFilePath(filename)
	if err != nil {
		fmt.Println(filepath)
		return err
	}

	cfgBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}

	cfg = &Config{}
	err = yaml.Unmarshal(cfgBytes, cfg)

	return err
}