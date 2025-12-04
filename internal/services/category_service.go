package services

import (
	"errors"

	"github.com/smaelmr/finance-api/internal/domain/contract/repository"
	"github.com/smaelmr/finance-api/internal/domain/entities"
)

type CategoryService struct {
	RepoManager repository.RepoManager
}

func NewCategoryService(repoManager repository.RepoManager) *CategoryService {
	return &CategoryService{
		RepoManager: repoManager,
	}
}

func (s *CategoryService) Add(category entities.Category) error {
	if category.Name == "" {
		return errors.New("nome não pode ser vazio")
	}

	if category.Type != "R" && category.Type != "D" {
		return errors.New("tipo deve ser 'R' (receita) ou 'D' (despesa)")
	}

	return s.RepoManager.Category().Add(category)
}

func (s *CategoryService) GetAll() ([]entities.Category, error) {
	categories, err := s.RepoManager.Category().GetAll()
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryService) GetByType(categoryType string) ([]entities.Category, error) {
	if categoryType != "R" && categoryType != "D" {
		return nil, errors.New("tipo deve ser 'R' (receita) ou 'D' (despesa)")
	}

	categories, err := s.RepoManager.Category().GetByType(categoryType)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryService) Delete(id int64) error {
	if id <= 0 {
		return errors.New("invalid id")
	}

	return s.RepoManager.Category().Delete(id)
}

func (s *CategoryService) Get(id int64) (*entities.Category, error) {
	if id <= 0 {
		return nil, errors.New("invalid id")
	}

	return s.RepoManager.Category().Get(id)
}
