package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	CompanyName string `mapstructure:"company_name"`
	ShortName   string `mapstructure:"short_name"`
	NIP         string `mapstructure:"nip"`
	REGON       string `mapstructure:"regon"`
	Email       string `mapstructure:"email"`
	Phone       string `mapstructure:"phone"`
}

func LoadConfig(cfgFile string) *Config {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("unable to decode into struct, %v", err))
	}

	return &config
}
