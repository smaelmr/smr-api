package controllers

import (
	"encoding/json"

	"net/http"

	"github.com/smaelmr/finance-api/api/commands"
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

func (c *PersonController) HandleCustomer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var person commands.PersonAdd
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.personService.Add(&person)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	case "GET":
		// Get a freight record by ID
		//id, err := strconv.Atoi(r.URL.Query().Get("id"))
		//if err != nil {
		//	w.WriteHeader(http.StatusBadRequest)
		//	return
		//}

		record, err := c.personService.GetCustomers()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(record)
		// Para DELETE e UPDATE, você pode adicionar os casos aqui.
	}
}

func (c *PersonController) HandleSupplier(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var person commands.PersonAdd
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.personService.Add(&person)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	case "GET":
		// Get a freight record by ID
		//id, err := strconv.Atoi(r.URL.Query().Get("id"))
		//if err != nil {
		//	w.WriteHeader(http.StatusBadRequest)
		//	return
		//}

		record, err := c.personService.GetSuppliers()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(record)
		// Para DELETE e UPDATE, você pode adicionar os casos aqui.
	}
}

func (c *PersonController) HandleDriver(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var person commands.PersonAdd
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = c.personService.Add(&person)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	case "GET":
		// Get a freight record by ID
		//id, err := strconv.Atoi(r.URL.Query().Get("id"))
		//if err != nil {
		//	w.WriteHeader(http.StatusBadRequest)
		//	return
		//}

		record, err := c.personService.GetDrivers()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(record)
		// Para DELETE e UPDATE, você pode adicionar os casos aqui.
	}
}
