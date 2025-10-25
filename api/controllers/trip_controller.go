package controllers

import (
	"encoding/json"
	"net/http"

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
		/*clienteId := r.URL.Query().Get("customerId")
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

		records, err := c.tripService.Filter(clienteIdPtr, motoristaIdPtr,
			dataInicialPtr, dataFinalPtr, placaCavaloPtr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}*/

		records, err := c.tripService.GetAll()
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
