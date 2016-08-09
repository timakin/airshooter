package config

import (
	"github.com/jinzhu/configor"
	"path/filepath"
)

type Config struct {
	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}
}

func Load() *Config {
	loaded := &Config{}
	absConfigPath, _ := filepath.Abs("./config.yml")
	configor.Load(loaded, absConfigPath)
	return loaded
}
