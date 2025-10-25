package services

import (
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"strconv"
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

func (c *FuelingService) ImportLinxDelPozo(file multipart.File, w http.ResponseWriter) ([]dto.FuelingImport, []string, bool) {
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
		fueling, err := c.convertToFueling(record)
		if err != nil {
			errors = append(errors, fmt.Sprintf("Erro ao processar registro: %s", err.Error()))
			continue
		}

		err = c.Add(&fueling)
		if err != nil {
			errors = append(errors, fmt.Sprintf("Erro ao salvar registro: %s", err.Error()))
			continue
		}
	}
	return importedRecords, errors, false
}

func (c *FuelingService) convertToFueling(record dto.FuelingImport) (entities.Fueling, error) {
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
	posto, err := c.personService.GetSupplierByCnpj(record.CnpjPosto)
	if err != nil {
		return entities.Fueling{}, fmt.Errorf("fornecedor não encontrado: %s", err)
	}

	// Buscar veiculo pela Placa
	veiculo, err := c.vehicleService.GetByPlate(record.Placa)
	if err != nil {
		return entities.Fueling{}, fmt.Errorf("Veiculo não encontrado: %s", err)
	}

	km, err := strconv.ParseInt(record.Hodometro, 10, 64)
	if err != nil {
		return entities.Fueling{}, fmt.Errorf("hodômetro inválido: %s", err)
	}

	cleanedValorTotal := c.cleanString(record.ValorTotal)
	valorTotalFloat, err := strconv.ParseFloat(cleanedValorTotal, 64)
	if err != nil {
		return entities.Fueling{}, fmt.Errorf("valor total inválido: %s", err)
	}

	quantidadeFloat, err := strconv.ParseFloat(strings.ReplaceAll(record.Quantidade, ",", "."), 64)
	if err != nil {
		return entities.Fueling{}, fmt.Errorf("quantidade inválida: %s", err)
	}

	return entities.Fueling{
		Data:            data,
		NumeroDocumento: record.NumeroCupom,
		Litros:          quantidadeFloat,
		ValorTotal:      valorTotalFloat,
		PostoId:         posto.Id,
		VeiculoId:       veiculo.Id,
		Km:              km,
	}, nil
}

func (c *FuelingService) cleanString(s string) string {
	// Remove todos os pontos e vírgulas da string
	s = strings.ReplaceAll(s, ".", "")
	s = strings.ReplaceAll(s, ",", "")
	return strings.TrimSpace(s)
}

func (c *FuelingService) parseFloat(s string) float64 {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0
	}
	return v
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

	//if dieselAdd.InicioViagem && dieselAdd.FinalViagem {
	//	err := errors.New("um abastecimento não pode ser de início e final de viagem ao mesmo tempo")
	//	return fmt.Errorf("não foi possível salvar o abastecimento: %w", err)
	//}

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

func (s *FuelingService) Delete(id int64) error {
	return s.RepoManager.Fueling().Delete(id)
}

/*func (s *FuelingService) Filter(fornecedorId *string, placa *string, dataInicial *string, dataFinal *string) ([]entities.Fueling, error) {

	filterParams := filter.NewFuelingFilterParams(fornecedorId, placa, dataInicial, dataFinal)

	dieselFilter, err := filterParams.ToFilter()
	if err != nil {
		return nil, err
	}

	return s.RepoManager.Fueling().Filter(*dieselFilter)
}*/
