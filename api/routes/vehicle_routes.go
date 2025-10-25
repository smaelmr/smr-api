package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/smaelmr/finance-api/api/controllers"
	"github.com/smaelmr/finance-api/internal/services"
)

func NewVehicleRoutes(service *services.VehicleService) chi.Router {
	router := chi.NewRouter()
	controller := controllers.NewVehicleController(service)

	router.Get("/", controller.GetAll)
	router.Post("/", controller.Create)
	router.Get("/{id}", controller.Get)
	router.Put("/{id}", controller.Update)
	router.Delete("/{id}", controller.Delete)

	return router
}
