package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
	"github.com/smaelmr/finance-api/api"
	"github.com/smaelmr/finance-api/config"
	"github.com/smaelmr/finance-api/internal/auth"
	"github.com/smaelmr/finance-api/internal/infrastructure/database"
	"github.com/smaelmr/finance-api/internal/infrastructure/database/repository"
)

func main() {
	conf := loadConfig()
	database := setupDatabase(conf)

	// Inicializar o serviço de autenticação
	jwtService := auth.NewJWTAuthService(conf.Auth.SecretKey)
	auth.InitAuthMiddleware(jwtService)

	// Inicializar o repositório
	repo := repository.NewRepo(database.DB)

	// Configurar o router
	router := api.SetupRouter(repo, jwtService)

	// Configurar o CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://smr-form.onrender.com", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Rodar o servidor com o middleware CORS
	handler := c.Handler(router)

	log.Println("Servidor rodando na porta 8088")
	log.Fatal(http.ListenAndServe(":8088", handler))
}

func loadConfig() *config.Config {
	log.Println("Carregando configuração...")

	// Primeiro tenta na raiz
	configFilePath := "./config.json"

	// Se não encontrar, tenta no diretório config
	if _, err := os.Stat(configFilePath); err != nil {
		configFilePath = "./../config"
	}

	log.Printf("Tentando carregar configuração de: %s", configFilePath)

	configFileName := "config"
	configFileExtension := "json"

	conf, err := config.LoadConfig(configFilePath, configFileName, configFileExtension)
	if err != nil {
		log.Fatalf("Erro ao carregar configuração: %v", err)
	}

	log.Printf("leu o config: %+s", conf.Auth.Teste)

	return conf
}

func setupDatabase(config *config.Config) *database.MySQLConnection {
	log.Println("Configurando banco de dados...")

	db, err := database.NewMySQLConnection(config)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	log.Println("Banco de dados configurado com sucesso")

	return db
}
