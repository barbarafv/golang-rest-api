package configuration

import (
	"log"

	"github.com/spf13/viper"
)

type dBconfig struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	DBUser   string `mapstructure:"DB_USER"`
	DBPass   string `mapstructure:"DB_PASS"`
}

var Config dBconfig

func init() {
	loadConfig("source/configuration/enviroments")
}

func loadConfig(path string) {
	viper.AddConfigPath(path)
	viper.SetConfigName("local")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Panic("<ReadInConfig> error to read in config")
	}
	err = viper.Unmarshal(&Config)

	if err != nil {
		log.Panic("<Unmarshal> error to Unmarshal")
	}
}
