package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/smaelmr/finance-api/api/controllers"
	"github.com/smaelmr/finance-api/internal/services"
)

func NewPaymentMethodRoutes(service *services.PaymentMethodService) chi.Router {
	router := chi.NewRouter()
	controller := controllers.NewPaymentMethodController(service)

	router.Get("/", controller.HandlePaymentMethod)
	router.Post("/", controller.HandlePaymentMethod)
	router.Get("/{id}", controller.HandlePaymentMethodById)
	router.Put("/{id}", controller.HandlePaymentMethodById)
	router.Delete("/{id}", controller.HandlePaymentMethodById)

	return router
}
