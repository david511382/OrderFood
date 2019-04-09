package config

import (
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

var (
	cfg *Config
)

func Get() *Config {
	return cfg
}

// ReadConfig read config from filepath
func ReadConfig(filename string) error {
	filepath, err := getFilePath(filename)
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

func getFilePath(filename string) (string, error) {
	if strings.HasPrefix(filename, "~") {
		u, err := user.Current()
		if err != nil {
			return filename, err
		}

		filename = strings.Replace(filename, "~", u.HomeDir, 1)
	} else {
		fn, err := filepath.Abs(filename)
		if err != nil {
			return filename, err
		}

		filename = fn
	}

	return filename, nil
}

func (c *Config) Domain() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}
