package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func ParseConfig() *Configuration {

	environment := ParseEnvironment(os.Getenv("ENVIRONMENT"))

	configFile := fmt.Sprintf("config/%s/config.json", environment)

	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var conf Configuration
	if err := viper.Unmarshal(&conf); err != nil {
		panic(err)
	}

	return &conf
}
