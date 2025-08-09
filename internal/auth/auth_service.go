package auth

import (
	"errors"
)

var (
	ErrInvalidToken = errors.New("token inválido")
	ErrExpiredToken = errors.New("token expirado")
)

type AuthService interface {
	ValidateToken(token string) error
}
