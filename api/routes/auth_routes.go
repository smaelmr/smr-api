package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/smaelmr/finance-api/api/controllers"
	"github.com/smaelmr/finance-api/internal/auth"
)

func NewAuthRoutes(jwtService *auth.JWTAuthService) chi.Router {
	router := chi.NewRouter()
	controller := controllers.NewLoginController(jwtService)

	router.Post("/login", controller.Login)

	return router
}
