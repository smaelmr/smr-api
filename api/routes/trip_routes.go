package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/smaelmr/finance-api/api/controllers"
	"github.com/smaelmr/finance-api/internal/services"
)

func NewTripRoutes(service *services.TripService) chi.Router {
	router := chi.NewRouter()
	controller := controllers.NewTripController(service)

	router.Get("/", controller.HandleTrip)
	router.Post("/", controller.HandleTrip)
	router.Put("/{id}", controller.HandleTrip)
	router.Delete("/{id}", controller.HandleTrip)

	return router
}
