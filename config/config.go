package config

import (
	"log"
	"os"

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
	// Configurar Viper para ler variáveis de ambiente
	viper.AutomaticEnv()

	// Mapear variáveis de ambiente para a configuração
	viper.BindEnv("Database.User", "DATABASE_USER")
	viper.BindEnv("Database.Pass", "DATABASE_PASS")
	viper.BindEnv("Database.Host", "DATABASE_HOST")
	viper.BindEnv("Database.Port", "DATABASE_PORT")
	viper.BindEnv("Database.Name", "DATABASE_NAME")
	viper.BindEnv("Auth.SecretKey", "AUTH_SECRETKEY")

	// Tentar carregar o arquivo de configuração
	viper.AddConfigPath(filepath)
	viper.SetConfigName(filename)
	viper.SetConfigType(extension)

	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("Arquivo de configuração não encontrado, usando variáveis de ambiente: %v", err)
	}

	// Se existirem variáveis de ambiente, elas terão prioridade
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	// Validar configuração
	if err := validateConfig(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func validateConfig(config *Config) error {
	// Verificar variáveis de ambiente primeiro
	if os.Getenv("DATABASE_USER") != "" {
		config.Database.User = os.Getenv("DATABASE_USER")
	}
	if os.Getenv("DATABASE_PASS") != "" {
		config.Database.Pass = os.Getenv("DATABASE_PASS")
	}
	if os.Getenv("DATABASE_HOST") != "" {
		config.Database.Host = os.Getenv("DATABASE_HOST")
	}
	if os.Getenv("DATABASE_PORT") != "" {
		if port := viper.GetInt("DATABASE_PORT"); port != 0 {
			config.Database.Port = port
		}
	}
	if os.Getenv("DATABASE_NAME") != "" {
		config.Database.Name = os.Getenv("DATABASE_NAME")
	}

	if os.Getenv("AUTH_SECRETKEY") != "" {
		config.Auth.SecretKey = os.Getenv("AUTH_SECRETKEY")
	}

	// Log das configurações (sem senhas)
	log.Printf("Configuração do banco de dados: Host=%s, Port=%d, User=%s, Database=%s",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Name,
	)

	return nil
}

func GetString(key string) string {
	// Verificar primeiro a variável de ambiente
	envValue := os.Getenv(key)
	if envValue != "" {
		return envValue
	}

	// Se não encontrar na variável de ambiente, usar o valor do Viper
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
