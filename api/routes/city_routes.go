package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/smaelmr/finance-api/api/controllers"
	"github.com/smaelmr/finance-api/internal/services"
)

func NewCityRoutes(service *services.CityService) chi.Router {
	router := chi.NewRouter()
	controller := controllers.NewCityController(service)

	router.Get("/", controller.HandleCity)
	router.Post("/", controller.HandleCity)
	router.Put("/{id}", controller.HandleCity)
	router.Delete("/{id}", controller.HandleCity)

	return router
}
