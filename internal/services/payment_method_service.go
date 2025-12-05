package services

import (
	"errors"

	"github.com/smaelmr/finance-api/internal/domain/contract/repository"
	"github.com/smaelmr/finance-api/internal/domain/entities"
)

type PaymentMethodService struct {
	RepoManager repository.RepoManager
}

func NewPaymentMethodService(repoManager repository.RepoManager) *PaymentMethodService {
	return &PaymentMethodService{
		RepoManager: repoManager,
	}
}

func (s *PaymentMethodService) Add(paymentMethod entities.PaymentMethod) error {
	if paymentMethod.Name == "" {
		return errors.New("description is required")
	}

	return s.RepoManager.PaymentMethod().Add(paymentMethod)
}

func (s *PaymentMethodService) GetAll() ([]entities.PaymentMethod, error) {
	paymentMethods, err := s.RepoManager.PaymentMethod().GetAll()
	if err != nil {
		return nil, err
	}

	return paymentMethods, nil
}

func (s *PaymentMethodService) Get(id int64) (*entities.PaymentMethod, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	return s.RepoManager.PaymentMethod().Get(id)
}

func (s *PaymentMethodService) Update(paymentMethod entities.PaymentMethod) error {
	if paymentMethod.Id <= 0 {
		return errors.New("invalid id")
	}

	if paymentMethod.Name == "" {
		return errors.New("description is required")
	}

	return s.RepoManager.PaymentMethod().Update(paymentMethod)
}

func (s *PaymentMethodService) Delete(id int64) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	return s.RepoManager.PaymentMethod().Delete(id)
}
