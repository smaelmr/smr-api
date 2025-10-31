package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/services"
)

type TripController struct {
	tripService *services.TripService
}

func NewTripController(tripService *services.TripService) *TripController {
	return &TripController{
		tripService: tripService,
	}
}

func (c *TripController) HandleTrip(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var trip entities.Trip
		err := json.NewDecoder(r.Body).Decode(&trip)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.tripService.Add(&trip)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	case "PUT":
		var trip entities.Trip
		err := json.NewDecoder(r.Body).Decode(&trip)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.tripService.Update(&trip)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	case "GET":
		monthStr := r.URL.Query().Get("month")
		yearStr := r.URL.Query().Get("year")

		var records []entities.Trip
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

			records, err = c.tripService.GetByMonthYear(month, year)
		} else {
			records, err = c.tripService.GetAll()
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
	}
}
