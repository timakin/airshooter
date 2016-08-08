package config

import (
	"github.com/jinzhu/configor"
)

var DB = struct {
	Name     string
	User     string `default:"root"`
	Password string `required:"true" env:"DBPassword"`
	Port     uint   `default:"3306"`
}{}
