package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/services"
)

type FreightController struct {
	freightService *services.FreightService
}

func NewFreightController(freightService *services.FreightService) *FreightController {
	return &FreightController{
		freightService: freightService,
	}
}

func (c *FreightController) HandleFreight(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var freight entities.Freight
		err := json.NewDecoder(r.Body).Decode(&freight)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.freightService.Add(&freight)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	case "PUT":
		var freight entities.Freight
		err := json.NewDecoder(r.Body).Decode(&freight)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.freightService.Update(&freight)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	case "GET":
		clienteId := r.URL.Query().Get("customerId")
		motoristaId := r.URL.Query().Get("driverId")
		dataInicial := r.URL.Query().Get("startDate")
		dataFinal := r.URL.Query().Get("endDate")
		placaCavalo := r.URL.Query().Get("truckPlate")

		var clienteIdPtr, motoristaIdPtr, dataInicialPtr, dataFinalPtr,
			placaCavaloPtr *string

		if clienteId != "" {
			clienteIdPtr = &clienteId
		}
		if motoristaId != "" {
			motoristaIdPtr = &motoristaId
		}
		if dataInicial != "" {
			dataInicialPtr = &dataInicial
		}
		if dataFinal != "" {
			dataFinalPtr = &dataFinal
		}
		if placaCavalo != "" {
			placaCavaloPtr = &placaCavalo
		}

		records, err := c.freightService.Filter(clienteIdPtr, motoristaIdPtr,
			dataInicialPtr, dataFinalPtr, placaCavaloPtr)
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
