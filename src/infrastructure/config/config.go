package config

import "github.com/kelseyhightower/envconfig"

const EnvDev = "development"

type Config struct {
	Environment string `required:"true" default:"development"`
	Port        int    `required:"true" default:"8080"`
	LogLevel    string `split_words:"true" default:"DEBUG"`
	MongoUri    string `split_words:"true" required:"true"`
}

func LoadConfig() Config {
	var config Config
	if err := envconfig.Process("", &config); err != nil {
		panic(err.Error())
	}
	return config
}
