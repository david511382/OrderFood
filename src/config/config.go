package config

import (
	"fmt"
	"strconv"
)

type Config struct {
	Server `yaml:"server"`
	MySQL  DbConfig `yaml:"mysql"`
	Txt    DbConfig `yaml:"txt"`
}

type Server struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type DbConfig struct {
	Domain   string `yaml:"domain"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Param    string `yaml:"param"`
}

var (
	cfg *Config
)

func Get() *Config {
	return cfg
}

func (c *Config) Domain() string {
	return c.Host + ":" + strconv.Itoa(c.Port)
}

// MysqlURL
func (d *DbConfig) MysqlURL() string {
	dsnFormat := "%s:%s@tcp(%s)/%s?%s"
	dsn := fmt.Sprintf(dsnFormat,
		d.Username,
		d.Password,
		d.Domain,
		d.Database,
		d.Param)
	return dsn
}
