package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/smaelmr/finance-api/internal/auth"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Aqui seria verificada a senha do banco de dados
	if creds.Username != "admin" || creds.Password != "senha123" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	tokenString, err := auth.GenerateJWT(creds.Username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//http.SetCookie(w, &http.Cookie{
	//	Name:  "token",
	//	Value: tokenString,
	//})

	// Enviar o token no corpo da resposta em formato JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(TokenResponse{Token: tokenString})
}
