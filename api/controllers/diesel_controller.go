package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/services"
)

type DieselController struct {
	dieselService *services.DieselService
}

func NewDieselController(dieselService *services.DieselService) *DieselController {
	return &DieselController{
		dieselService: dieselService,
	}
}

func (c *DieselController) HandleDiesel(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var diesel entities.Diesel
		err := json.NewDecoder(r.Body).Decode(&diesel)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.dieselService.Add(&diesel)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	case "PUT":
		var dieselUpdate entities.Diesel
		err := json.NewDecoder(r.Body).Decode(&dieselUpdate)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.dieselService.Update(&dieselUpdate)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	case "GET":
		fornecedorId := r.URL.Query().Get("supplierId")
		placa := r.URL.Query().Get("plate")
		dataInicial := r.URL.Query().Get("startDate")
		dataFinal := r.URL.Query().Get("endDate")

		var fornecedorIdPtr, placaPtr, dataInicialPtr, dataFinalPtr *string

		if fornecedorId != "" {
			fornecedorIdPtr = &fornecedorId
		}
		if placa != "" {
			placaPtr = &placa
		}
		if dataInicial != "" {
			dataInicialPtr = &dataInicial
		}
		if dataFinal != "" {
			dataFinalPtr = &dataFinal
		}

		records, err := c.dieselService.Filter(fornecedorIdPtr, placaPtr, dataInicialPtr, dataFinalPtr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(records) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(records)
	}
}
