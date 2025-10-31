package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/smaelmr/finance-api/api/controllers"
	"github.com/smaelmr/finance-api/internal/services"
)

func NewFuelingRoutes(dieselService *services.FuelingService, personService *services.PersonService) chi.Router {
	router := chi.NewRouter()
	controller := controllers.NewFuelingController(dieselService, personService)

	router.Get("/", controller.HandleFueling)
	router.Post("/", controller.HandleFueling)
	router.Put("/{id}", controller.HandleFueling)
	router.Delete("/{id}", controller.HandleFueling)

	router.Post("/import-russi", controller.HandleImportLinxDelPozo)
	router.Get("/consumo", controller.HandleGetConsumption)

	return router
}
