package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/smaelmr/finance-api/internal/domain/entities"
	"github.com/smaelmr/finance-api/internal/services"
)

type PersonController struct {
	personService *services.PersonService
}

func NewPersonController(personService *services.PersonService) *PersonController {
	return &PersonController{
		personService: personService,
	}
}

func (c *PersonController) HandleClient(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var person entities.Person
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.personService.AddClient(&person)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	case "GET":
		record, err := c.personService.GetClients()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(record)
	case "PUT":
		id := chi.URLParam(r, "id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var person entities.Person
		err = json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		person.Id = idInt
		err = c.personService.UpdateClient(&person)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	case "DELETE":
		id := chi.URLParam(r, "id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.personService.DeleteClient(idInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (c *PersonController) HandleSupplier(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var person entities.Person
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.personService.AddSupplier(&person)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	case "GET":
		record, err := c.personService.GetSuppliers()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(record)
	case "PUT":
		id := chi.URLParam(r, "id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var person entities.Person
		err = json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		person.Id = idInt
		err = c.personService.UpdateSupplier(&person)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	case "DELETE":
		id := chi.URLParam(r, "id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.personService.DeleteSupplier(idInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (c *PersonController) HandleGasStation(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var person entities.Person
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.personService.AddGasStation(&person)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	case "GET":
		record, err := c.personService.GetGasStation()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(record)
	case "PUT":
		id := chi.URLParam(r, "id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var person entities.Person
		err = json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		person.Id = idInt
		err = c.personService.UpdateGasStation(&person)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	case "DELETE":
		id := chi.URLParam(r, "id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.personService.DeleteGasStation(idInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func (c *PersonController) HandleDriver(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var driver entities.Driver
		err := json.NewDecoder(r.Body).Decode(&driver)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.personService.AddDriver(&driver)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	case "GET":
		record, err := c.personService.GetDrivers()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(record)
	case "PUT":
		id := chi.URLParam(r, "id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var driver entities.Driver
		err = json.NewDecoder(r.Body).Decode(&driver)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		driver.Id = idInt
		err = c.personService.UpdateDriver(&driver)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	case "DELETE":
		id := chi.URLParam(r, "id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		idInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.personService.DeleteDriver(idInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
