package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/smaelmr/finance-api/api/routes"
	"github.com/smaelmr/finance-api/internal/auth"
	"github.com/smaelmr/finance-api/internal/infrastructure/database/repository"
	"github.com/smaelmr/finance-api/internal/services"
)

func SetupRouter(repo *repository.Repo, jwtService *auth.JWTAuthService) *chi.Mux {
	r := chi.NewRouter()

	// Middleware global
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	// Inicializar o middleware de autenticação
	auth.InitAuthMiddleware(jwtService)

	// Rotas públicas
	r.Group(func(r chi.Router) {
		r.Mount("/api/v1/auth", routes.NewAuthRoutes(jwtService))
	})

	// Rotas protegidas
	r.Group(func(r chi.Router) {
		//r.Use(auth.AuthMiddleware)

		// Serviços
		personService := services.NewPersonService(repo)
		vehicleService := services.NewVehicleService(repo)
		dieselService := services.NewFuelingService(repo, personService, vehicleService)
		cityService := services.NewCityService(repo)
		tripService := services.NewTripService(repo)
		categoryService := services.NewCategoryService(repo)

		// Montar rotas
		r.Mount("/api/v1/person", routes.NewPersonRoutes(personService))
		r.Mount("/api/v1/fueling", routes.NewFuelingRoutes(dieselService, personService))
		r.Mount("/api/v1/city", routes.NewCityRoutes(cityService))
		r.Mount("/api/v1/trip", routes.NewTripRoutes(tripService))
		r.Mount("/api/v1/vehicle", routes.NewVehicleRoutes(vehicleService))
		r.Mount("/api/v1/category", routes.NewCategoryRoutes(categoryService))
	})

	return r
}
