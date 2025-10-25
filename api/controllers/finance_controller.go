package controllers

import (
	"encoding/json"

	"net/http"

	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/services"
)

type FinanceController struct {
	financeService *services.FinanceService
}

func NewFinanceController(financeService *services.FinanceService) *FinanceController {
	return &FinanceController{
		financeService: financeService,
	}
}

func (c *FinanceController) HandleFinance(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		// Create a new finance record
		var record entities.Finance
		err := json.NewDecoder(r.Body).Decode(&record)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.financeService.Add(record)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	case "GET":
		// Get a finance record by ID
		//id, err := strconv.Atoi(r.URL.Query().Get("id"))
		//if err != nil {
		//	w.WriteHeader(http.StatusBadRequest)
		//	return
		//}

		record, err := c.financeService.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(record)
		// Para DELETE e UPDATE, você pode adicionar os casos aqui.
	}
}
