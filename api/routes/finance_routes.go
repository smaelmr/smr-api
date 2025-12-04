package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/smaelmr/finance-api/api/controllers"
	"github.com/smaelmr/finance-api/internal/services"
)

func NewFinanceRoutes(service *services.FinanceService) chi.Router {
	router := chi.NewRouter()
	controller := controllers.NewFinanceController(service)

	router.Get("/receipts", controller.HandleReceipts)
	router.Get("/payments", controller.HandlePayments)
	router.Post("/", controller.HandleFinance)
	router.Put("/{id}/payment", controller.HandlePayment)
	router.Delete("/{id}", controller.HandleFinance)
	router.Get("/{id}", controller.HandleFinance)

	return router
}
