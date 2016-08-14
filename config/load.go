package config

import (
	"github.com/jinzhu/configor"
	"path/filepath"
)

type Config struct {
	AppPort string `default:"4000"`

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}
}

func Load() *Config {
	loaded := &Config{}
	configPath, _ := filepath.Abs("./config.yml")
	configor.Load(loaded, configPath)
	return loaded
}
