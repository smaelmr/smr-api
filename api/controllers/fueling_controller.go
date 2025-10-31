package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	entities "github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/services"
)

type FuelingController struct {
	fuelingService *services.FuelingService
	personService  *services.PersonService
}

func NewFuelingController(fuelingService *services.FuelingService, personService *services.PersonService) *FuelingController {
	return &FuelingController{
		fuelingService: fuelingService,
		personService:  personService,
	}
}

func (c *FuelingController) HandleFueling(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var fueling entities.Fueling
		err := json.NewDecoder(r.Body).Decode(&fueling)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.fuelingService.Add(&fueling)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	case "PUT":
		var fuelingUpdate entities.Fueling
		err := json.NewDecoder(r.Body).Decode(&fuelingUpdate)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.fuelingService.Update(&fuelingUpdate)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	case "GET":
		monthStr := r.URL.Query().Get("month")
		yearStr := r.URL.Query().Get("year")

		var records []entities.Fueling
		var err error

		if monthStr != "" && yearStr != "" {
			month, err := strconv.Atoi(monthStr)
			if err != nil {
				http.Error(w, "Mês inválido", http.StatusBadRequest)
				return
			}

			year, err := strconv.Atoi(yearStr)
			if err != nil {
				http.Error(w, "Ano inválido", http.StatusBadRequest)
				return
			}

			if month < 1 || month > 12 {
				http.Error(w, "Mês deve estar entre 1 e 12", http.StatusBadRequest)
				return
			}

			records, err = c.fuelingService.GetByMonthYear(month, year)
		} else {
			records, err = c.fuelingService.GetAll()
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(records) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(records)
	case "DELETE":
		idStr := chi.URLParam(r, "id")

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = c.fuelingService.Delete(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (c *FuelingController) HandleGetConsumption(w http.ResponseWriter, r *http.Request) {
	monthStr := r.URL.Query().Get("month")
	yearStr := r.URL.Query().Get("year")

	if monthStr != "" && yearStr != "" {
		month, err := strconv.Atoi(monthStr)
		if err != nil {
			http.Error(w, "Mês inválido", http.StatusBadRequest)
			return
		}

		year, err := strconv.Atoi(yearStr)
		if err != nil {
			http.Error(w, "Ano inválido", http.StatusBadRequest)
			return
		}

		if month < 1 || month > 12 {
			http.Error(w, "Mês deve estar entre 1 e 12", http.StatusBadRequest)
			return
		}

		consumos, err := c.fuelingService.GetFuelConsumption(month, year)
		if err != nil {
			http.Error(w, "Erro ao calcular consumo: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(consumos)

	} else {
		http.Error(w, "Mês e ano são obrigatórios", http.StatusBadRequest)
		return
	}
}

func (c *FuelingController) HandleImportLinxDelPozo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Receber o arquivo
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Erro ao receber arquivo: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Ler o arquivo Excel
	importedRecords, errors, shouldReturn := c.fuelingService.ImportLinxDelPozo(file, handler, w)
	if shouldReturn {
		return
	}

	response := map[string]interface{}{
		"success":          len(errors) == 0,
		"errors":           errors,
		"recordsProcessed": len(importedRecords),
		"recordsFailed":    len(errors),
	}

	json.NewEncoder(w).Encode(response)
}
