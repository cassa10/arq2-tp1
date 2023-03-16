package config

import "github.com/kelseyhightower/envconfig"

type Config struct {
	Environment string `required:"true" default:"development"`
	Port        int    `required:"true" default:"8888"`
	LogLevel    string `split_words:"true" default:"DEBUG"`
}

func LoadConfig() Config {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		panic(err.Error())
	}
	return config
}
