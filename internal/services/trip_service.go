package services

import (
	"time"

	"github.com/smaelmr/finance-api/internal/domain/contract/repository"
	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/domain/entities/filter"
)

type TripService struct {
	RepoManager repository.RepoManager
}

func NewTripService(repoManager repository.RepoManager) *TripService {
	return &TripService{
		RepoManager: repoManager,
	}
}

func (s *TripService) Add(tripAdd *entities.Trip) error {
	// Adiciona o frete e obtém o ID
	tripId, err := s.RepoManager.Trip().Add(*tripAdd)
	if err != nil {
		return err
	}

	// Cria automaticamente um lançamento a receber
	origemId := tripId
	finance := entities.Finance{
		PessoaId:        tripAdd.ClienteId,
		CategoriaId:     3, // categoria de "Frete"
		OrigemId:        &origemId,
		Origem:          "FRETE",
		Valor:           float64(tripAdd.ValorFrete),
		ValorParcela:    float64(tripAdd.ValorFrete),
		NumeroParcela:   1,
		TotalParcelas:   1,
		NumeroDocumento: tripAdd.NumeroDocumento,
		DataCompetencia: tripAdd.DataColeta,
		DataVencimento:  tripAdd.DataColeta,
		Observacao:      "Lançamento automático de frete",
	}

	err = s.RepoManager.Finance().Add(finance)
	if err != nil {
		// Se falhar ao criar o lançamento, não impede a criação do frete
		// mas registra o erro
		// TODO: Implementar log de erro
	}

	return nil
}

func (s *TripService) GetAll() ([]entities.Trip, error) {
	records, err := s.RepoManager.Trip().GetAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (s *TripService) Update(tripUpdate *entities.Trip) error {
	return s.RepoManager.Trip().Update(*tripUpdate)
}

func (s *TripService) GetByMonthYear(month, year int) ([]entities.Trip, error) {
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	endDate := startDate.AddDate(0, 1, 0).Add(-time.Second)

	return s.RepoManager.Trip().GetByDateRange(startDate, endDate)
}

func (s *TripService) Filter(clienteId, motoristaId, dataInicial, dataFinal, cavaloPlaca *string) ([]entities.Trip, error) {
	filterParams := filter.NewTripFilterParams(clienteId, motoristaId, dataInicial, dataFinal, cavaloPlaca)

	tripFilter, err := filterParams.ToFilter()
	if err != nil {
		return nil, err
	}

	records, err := s.RepoManager.Trip().Filter(*tripFilter)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (s *TripService) Delete(id int64) error {
	return s.RepoManager.Trip().Delete(id)
}
