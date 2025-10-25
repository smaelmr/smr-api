package auth

import (
	"net/http"
	"strings"
)

var authService AuthService

func InitAuthMiddleware(service AuthService) {
	authService = service
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Pular verificação para OPTIONS (necessário para CORS)
		if r.Method == "OPTIONS" {
			next.ServeHTTP(w, r)
			return
		}

		// Pegar o header Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		// Verificar se é um Bearer token
		splitToken := strings.Split(authHeader, "Bearer ")
		if len(splitToken) != 2 {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		token := splitToken[1]

		if err := authService.ValidateToken(token); err != nil {
			switch err {
			case ErrExpiredToken:
				http.Error(w, "Token expirado", http.StatusUnauthorized)
			case ErrInvalidToken:
				http.Error(w, "Token inválido", http.StatusUnauthorized)
			default:
				http.Error(w, "Erro na autenticação", http.StatusUnauthorized)
			}
			return
		}

		// Token válido, continuar com a requisição
		next.ServeHTTP(w, r)
	})
}
