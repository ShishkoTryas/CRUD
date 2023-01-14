package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Config struct {
	HttpAddr string `default:":8080"`
}

func NewConfig() *Config {
	var s Config
	err := envconfig.Process("service", &s)
	if err != nil {
		log.Fatal(err)
	}
	return &s
}
