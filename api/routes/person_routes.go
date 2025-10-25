package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/smaelmr/finance-api/api/controllers"
	"github.com/smaelmr/finance-api/internal/services"
)

func NewPersonRoutes(service *services.PersonService) chi.Router {
	router := chi.NewRouter()
	controller := controllers.NewPersonController(service)

	router.Get("/client", controller.HandleClient)
	router.Post("/client", controller.HandleClient)
	router.Put("/client/{id}", controller.HandleClient)
	router.Delete("/client/{id}", controller.HandleClient)

	router.Get("/supplier", controller.HandleSupplier)
	router.Post("/supplier", controller.HandleSupplier)
	router.Put("/supplier/{id}", controller.HandleSupplier)
	router.Delete("/supplier/{id}", controller.HandleSupplier)

	router.Get("/gas-station", controller.HandleGasStation)
	router.Post("/gas-station", controller.HandleGasStation)
	router.Put("/gas-station/{id}", controller.HandleGasStation)
	router.Delete("/gas-station/{id}", controller.HandleGasStation)

	router.Get("/driver", controller.HandleDriver)
	router.Post("/driver", controller.HandleDriver)
	router.Put("/driver/{id}", controller.HandleDriver)
	router.Delete("/driver/{id}", controller.HandleDriver)

	return router
}
