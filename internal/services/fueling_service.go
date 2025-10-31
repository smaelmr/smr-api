package services

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/smaelmr/finance-api/internal/domain/contract/repository"
	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/domain/entities/dto"
	"github.com/xuri/excelize/v2"
)

type FuelingService struct {
	RepoManager    repository.RepoManager
	personService  *PersonService
	vehicleService *VehicleService
}

func NewFuelingService(
	repoManager repository.RepoManager,
	personService *PersonService,
	vehicleService *VehicleService) *FuelingService {
	return &FuelingService{
		RepoManager:    repoManager,
		personService:  personService,
		vehicleService: vehicleService,
	}
}

func (c *FuelingService) ImportLinxDelPozo(file multipart.File, handler *multipart.FileHeader, w http.ResponseWriter) ([]dto.FuelingImport, []string, bool) {
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		http.Error(w, "Erro ao ler arquivo Excel: "+err.Error(), http.StatusBadRequest)
		return nil, nil, true
	}

	// Pegar a primeira planilha
	sheetName := xlsx.GetSheetName(0)
	rows, err := xlsx.GetRows(sheetName)
	if err != nil {
		http.Error(w, "Erro ao ler linhas do Excel: "+err.Error(), http.StatusInternalServerError)
		return nil, nil, true
	}

	var importedRecords []dto.FuelingImport
	var errors []string

	// Pular o cabeçalho
	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) < 10 {
			errors = append(errors, fmt.Sprintf("Linha %d: número de colunas inválido", i+1))
			continue
		}

		record := dto.FuelingImport{
			DataTransacao: row[0],
			NumeroCupom:   row[1],
			CnpjPosto:     row[2],
			Placa:         row[3],
			CpfMotorista:  row[4],
			Hodometro:     row[5],
			Produto:       row[6],
			ValorUnitario: row[7],
			Quantidade:    row[8],
			ValorTotal:    row[9],
		}

		// Validações básicas
		if !c.isValidDate(record.DataTransacao) {
			errors = append(errors, fmt.Sprintf("Linha %d: data inválida", i+1))
			continue
		}

		if !c.isValidCNPJ(record.CnpjPosto) {
			errors = append(errors, fmt.Sprintf("Linha %d: CNPJ inválido", i+1))
			continue
		}

		importedRecords = append(importedRecords, record)
	}

	if len(errors) > 0 {
		response := map[string]interface{}{
			"success": false,
			"errors":  errors,
		}
		json.NewEncoder(w).Encode(response)
		return nil, nil, true
	}

	// Processar os registros válidos
	for _, record := range importedRecords {

		fueling, err := c.convertToFueling(record, handler.Filename)
		if err != nil {
			errors = append(errors, fmt.Sprintf("Erro ao processar registro: %s", err.Error()))
			continue
		}

		if (fueling.TipoCombustivel == "Diesel_S10" || fueling.TipoCombustivel == "Diesel_S500" || fueling.TipoCombustivel == "Arla") == false {
			err = c.Add(&fueling)
			if err != nil {
				errors = append(errors, fmt.Sprintf("Erro ao salvar registro: %s", err.Error()))
				continue
			}
		}
	}

	return importedRecords, errors, false
}

func (c *FuelingService) convertToFueling(record dto.FuelingImport, fileName string) (entities.Fueling, error) {
	// Tenta primeiro o formato com hora e minuto
	data, err := time.Parse("02/01/2006 15:04", record.DataTransacao)
	if err != nil {
		// Se falhar, tenta o formato apenas com data
		data, err = time.Parse("02/01/2006", record.DataTransacao)
		if err != nil {
			return entities.Fueling{}, fmt.Errorf("data inválida, use o formato dd/MM/yyyy HH:mm ou dd/MM/yyyy: %s", err)
		}
	}

	// Buscar fornecedor pelo CNPJ
	posto, err := c.personService.GetGasStationByCnpj(record.CnpjPosto)
	if err != nil {
		return entities.Fueling{}, fmt.Errorf("fornecedor não encontrado: %s", err)
	}

	// Buscar veiculo pela Placa
	if record.Placa == "ISM5693" {
		record.Placa = "ISM5G93"
	}

	veiculo, err := c.vehicleService.GetByPlate(record.Placa)
	if err != nil {
		return entities.Fueling{}, fmt.Errorf("Veiculo não encontrado: %s", err)
	}

	fuel := entities.Fueling{
		Data:            data,
		NumeroDocumento: record.NumeroCupom,
		Litros:          record.QuantidadeFloat64(),
		ValorTotal:      record.ValorTotalFloat64(),
		PostoId:         posto.Id,
		VeiculoId:       veiculo.Id,
		Km:              record.HodometroInt64(),
	}

	// Mapeia o tipo de combustível baseado no nome do arquivo
	fileNameLower := strings.ToLower(fileName)
	if strings.Contains(fileNameLower, "russi") {
		fuel.TipoCombustivel = record.ProdutoMappedRussi()
	} else if strings.Contains(fileNameLower, "graal") {
		fuel.TipoCombustivel = record.ProdutoMappedGraal()
	} else {
		return entities.Fueling{}, fmt.Errorf("formato de arquivo não reconhecido: %s. Use arquivos com 'Russi' ou 'Graal' no nome", fileName)
	}

	return fuel, nil
}

func (c *FuelingService) isValidDate(date string) bool {
	// Tenta primeiro o formato com hora e minuto
	_, err := time.Parse("02/01/2006 15:04", date)
	if err == nil {
		return true
	}

	// Se falhar, tenta o formato apenas com data
	_, err = time.Parse("02/01/2006", date)
	return err == nil
}

func (c *FuelingService) isValidCNPJ(cnpj string) bool {
	// Implementar validação de CNPJ
	return len(cnpj) == 14
}

func (s *FuelingService) Add(dieselAdd *entities.Fueling) error {

	return s.RepoManager.Fueling().Add(*dieselAdd)
}

func (s *FuelingService) GetAll() ([]entities.Fueling, error) {
	records, err := s.RepoManager.Fueling().GetAll()
	if err != nil {
		return nil, err
	}

	var dieselList []entities.Fueling
	dieselList = append(dieselList, records...)

	return dieselList, nil
}

func (s *FuelingService) Update(dieselUpdate *entities.Fueling) error {
	return s.RepoManager.Fueling().Update(*dieselUpdate)
}

func (s *FuelingService) GetByMonthYear(month, year int) ([]entities.Fueling, error) {
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	endDate := startDate.AddDate(0, 1, 0).Add(-time.Second)

	return s.RepoManager.Fueling().GetByDateRange(startDate, endDate)
}

func (s *FuelingService) Delete(id int64) error {
	return s.RepoManager.Fueling().Delete(id)
}

// GetFuelConsumption retorna o consumo médio de combustível por veículo no período
func (s *FuelingService) GetFuelConsumption(month, year int) ([]dto.FuelingConsumption, error) {
	startDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local)
	endDate := startDate.AddDate(0, 1, 0).Add(-time.Second)

	return s.RepoManager.Fueling().GetFuelConsumption(startDate, endDate)
}
