package models

import "time"

type Order struct {
    ID       int64     `json:"id"`
    PetID    int64     `json:"petId"`
    Quantity int       `json:"quantity"`
    ShipDate time.Time `json:"shipDate"`
}