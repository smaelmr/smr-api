package config

import (
	"github.com/spf13/viper"
)

func GetEnvOrDefault(key, defaultValue string) string {
	value := GetString(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func LoadConfig(filepath, filename, extension string) (*Config, error) {
	viper.AddConfigPath(filepath)
	viper.SetConfigName(filename)
	viper.SetConfigType(extension)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func GetString(key string) string {
	value := viper.GetString(key)
	if value == "" {
		return ""
	}
	return value
}

func GetInt(key string) int {
	value := viper.GetInt(key)
	if value == 0 {
		return 0
	}
	return value
}
