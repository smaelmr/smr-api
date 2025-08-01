package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
	"github.com/smaelmr/finance-api/api/controllers"
	"github.com/smaelmr/finance-api/config"
	"github.com/smaelmr/finance-api/internal/infrastructure/database"
	"github.com/smaelmr/finance-api/internal/infrastructure/database/repository"
	"github.com/smaelmr/finance-api/internal/services"
)

func main() {

	conf := loadConfig()
	database := setupDatabase(conf)

	repo := repository.NewRepo(database.DB)

	dieselService := services.NewDieselService(repo)
	personService := services.NewPersonService(repo)
	cityService := services.NewCityService(repo)
	freightService := services.NewFreightService(repo)

	dieselController := controllers.NewDieselController(dieselService)
	personController := controllers.NewPersonController(personService)
	cityController := controllers.NewCityController(cityService)
	freightController := controllers.NewFreightController(freightService)

	// Definir os endpoints
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/login", controllers.Login)
	//mux.HandleFunc("/api/v1/finance", controllers.HandleFinance)
	mux.HandleFunc("/api/v1/customer", personController.HandleCustomer)
	mux.HandleFunc("/api/v1/supplier", personController.HandleSupplier)

	mux.HandleFunc("/api/v1/freight", freightController.HandleFreight)

	mux.HandleFunc("/api/v1/diesel", dieselController.HandleDiesel)
	mux.HandleFunc("/api/v1/driver", personController.HandleDriver)
	mux.HandleFunc("/api/v1/city", cityController.HandleCity)

	// Configurar o CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "https://smr-api.onrender.com"}, // Origem permitida
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	})

	// Rodar o servidor com o middleware CORS
	handler := c.Handler(mux)

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
