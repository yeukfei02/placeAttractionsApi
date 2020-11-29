package model

import "github.com/kamva/mgm/v3"

// PlaceAttractions struct
type PlaceAttractions struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string   `json:"name" bson:"name"`
	Rating           float64  `json:"rating"`
	Location         Location `json:"location"`
}

// Location struct
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
