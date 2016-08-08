package config

import (
	"github.com/jinzhu/configor"
	"path/filepath"
)

func Load() interface{} {
	var Config = struct {
		DB struct {
			Name     string
			User     string `default:"root"`
			Password string `required:"true" env:"DBPassword"`
			Port     uint   `default:"3306"`
		}
	}{}

	absConfigPath, _ := filepath.Abs("./config.yml")
	configor.Load(&Config, absConfigPath)
	return &Config
}
