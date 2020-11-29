package model

import "github.com/kamva/mgm/v3"

// PlaceAttractions struct
type PlaceAttractions struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string   `json:"name" bson:"name"`
	Rating           float64  `json:"rating" bson:"rating"`
	Location         Location `json:"location" bson:"location"`
}

// Location struct
type Location struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lng float64 `json:"lng" bson:"lng"`
}
