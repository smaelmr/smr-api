package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/domain/entities/dto"
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

func (c *FinanceController) HandleReceipts(w http.ResponseWriter, r *http.Request) {
	monthStr := r.URL.Query().Get("month")
	yearStr := r.URL.Query().Get("year")

	if monthStr == "" || yearStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "month and year are required"})
		return
	}

	month := 0
	year := 0
	_, err := fmt.Sscanf(monthStr, "%d", &month)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid month"})
		return
	}

	_, err = fmt.Sscanf(yearStr, "%d", &year)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid year"})
		return
	}

	record, err := c.financeService.GetReceipts(month, year)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
}

func (c *FinanceController) HandlePayments(w http.ResponseWriter, r *http.Request) {
	monthStr := r.URL.Query().Get("month")
	yearStr := r.URL.Query().Get("year")

	if monthStr == "" || yearStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "month and year are required"})
		return
	}

	month := 0
	year := 0
	_, err := fmt.Sscanf(monthStr, "%d", &month)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid month"})
		return
	}

	_, err = fmt.Sscanf(yearStr, "%d", &year)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid year"})
		return
	}

	record, err := c.financeService.GetPayments(month, year)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(record)
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

		categoryType := r.URL.Query().Get("type")
		monthStr := r.URL.Query().Get("month")
		yearStr := r.URL.Query().Get("year")

		if categoryType == "" || monthStr == "" || yearStr == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "type, month and year are required"})
			return
		}

		month := 0
		year := 0
		_, err := fmt.Sscanf(monthStr, "%d", &month)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid month"})
			return
		}

		_, err = fmt.Sscanf(yearStr, "%d", &year)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "invalid year"})
			return
		}

		record, err := c.financeService.GetAll(categoryType, month, year)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(record)
		// Para DELETE e UPDATE, você pode adicionar os casos aqui.
	}
}

func (c *FinanceController) HandlePayment(w http.ResponseWriter, r *http.Request) {
	// Apenas PUT é permitido
	if r.Method != "PUT" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(map[string]string{"error": "Method not allowed"})
		return
	}

	// Obter o ID do lançamento da URL
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid ID"})
		return
	}

	// Decodificar o body da requisição
	var paymentReq dto.PaymentRequest
	err = json.NewDecoder(r.Body).Decode(&paymentReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request body"})
		return
	}

	// Processar o pagamento
	err = c.financeService.ProcessPayment(
		id,
		paymentReq.ValorPago,
		paymentReq.DataRealizacao,
		paymentReq.FormaPagamentoId,
		paymentReq.LancarDiferenca,
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Payment processed successfully"})
}
