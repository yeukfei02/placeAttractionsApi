package place

import (
	"github.com/kamva/mgm/v3"
	"github.com/yeukfei02/placeAttractionsApi/src/helpers"
	placeAttractionsModel "github.com/yeukfei02/placeAttractionsApi/src/model/placeAttractions"
	"go.mongodb.org/mongo-driver/bson"
)

// CreatePlaceAttractions func
func CreatePlaceAttractions(name string, rating float64, lat float64, lng float64) {
	placeAttractions := &placeAttractionsModel.PlaceAttractions{
		Name:   name,
		Rating: rating,
		Location: placeAttractionsModel.Location{
			Lat: lat,
			Lng: lng,
		},
	}

	err := mgm.Coll(placeAttractions).Create(placeAttractions)
	helpers.CheckErr(err)
}

// GetPlaceAttractions func
func GetPlaceAttractions() []placeAttractionsModel.PlaceAttractions {
	result := []placeAttractionsModel.PlaceAttractions{}

	err := mgm.Coll(&placeAttractionsModel.PlaceAttractions{}).SimpleFind(&result, bson.M{})
	helpers.CheckErr(err)

	return result
}

// GetPlaceAttractionsByID func
func GetPlaceAttractionsByID(id string) *placeAttractionsModel.PlaceAttractions {
	placeAttractions := &placeAttractionsModel.PlaceAttractions{}

	err := mgm.Coll(placeAttractions).FindByID(id, placeAttractions)
	helpers.CheckErr(err)

	return placeAttractions
}

// UpdatePlaceAttractionsByID func
func UpdatePlaceAttractionsByID(id string, name string, rating float64, lat float64, lng float64) {
	placeAttractions := &placeAttractionsModel.PlaceAttractions{}

	err := mgm.Coll(placeAttractions).FindByID(id, placeAttractions)
	helpers.CheckErr(err)

	placeAttractions.Name = name
	placeAttractions.Rating = rating
	placeAttractions.Location.Lat = lat
	placeAttractions.Location.Lng = lng

	updateErr := mgm.Coll(placeAttractions).Update(placeAttractions)
	helpers.CheckErr(updateErr)
}

// DeletePlaceAttractionsByID func
func DeletePlaceAttractionsByID(id string) {
	placeAttractions := &placeAttractionsModel.PlaceAttractions{}

	err := mgm.Coll(placeAttractions).FindByID(id, placeAttractions)
	helpers.CheckErr(err)

	deleteErr := mgm.Coll(placeAttractions).Delete(placeAttractions)
	helpers.CheckErr(deleteErr)
}
