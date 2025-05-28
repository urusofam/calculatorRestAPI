package config

import "github.com/spf13/viper"

type Config struct {
	Database struct {
		Host string
		Port int
		User string
		Pass string
		Name string
	}
	Server struct {
		Host string
		Port int
	}
}

func LoadConfig() (*Config, error) {
	config := &Config{}

	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
