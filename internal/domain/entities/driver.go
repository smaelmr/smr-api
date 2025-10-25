package entities

import "time"

type Driver struct {
	Person
	CnhCategory string    `json:"cnhCategory"`
	CnhValidity time.Time `json:"cnhValidity"`
	Observation string    `json:"observation"`
}
