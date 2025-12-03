package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/smaelmr/finance-api/api/controllers"
	"github.com/smaelmr/finance-api/internal/services"
)

func NewCategoryRoutes(service *services.CategoryService) chi.Router {
	router := chi.NewRouter()
	controller := controllers.NewCategoryController(service)

	router.Get("/", controller.HandleCategory)
	router.Post("/", controller.HandleCategory)
	router.Delete("/{id}", controller.HandleCategoryById)
	router.Get("/{id}", controller.HandleCategoryById)

	return router
}
