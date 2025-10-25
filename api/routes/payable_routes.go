package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/smaelmr/finance-api/internal/services"
)

func NewPayableRoutes(service *services.TripService) chi.Router {
	router := chi.NewRouter()
	//controller := controllers.NewTripController(service)

	//router.Get("/", controller.HandlePayable)
	//router.Post("/", controller.HandlePayable)
	//router.Put("/{id}", controller.HandlePayable)
	//router.Delete("/{id}", controller.HandlePayable)

	return router
}
