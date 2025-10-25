package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/services"
)

type CityController struct {
	cityService *services.CityService
}

func NewCityController(cityService *services.CityService) *CityController {
	return &CityController{
		cityService: cityService,
	}
}

func (c *CityController) HandleCity(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var city entities.City
		err := json.NewDecoder(r.Body).Decode(&city)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.cityService.Add(city)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	case "GET":
		// Get a trip record by ID
		//id, err := strconv.Atoi(r.URL.Query().Get("id"))
		//if err != nil {
		//	w.WriteHeader(http.StatusBadRequest)
		//	return
		//}

		record, err := c.cityService.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(record)
		// Para DELETE e UPDATE, você pode adicionar os casos aqui.
	}
}
