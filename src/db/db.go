package db

import (
	"fmt"

	"github.com/kamva/mgm/v3"
	"github.com/yeukfei02/placeAttractionsApi/src/common"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB func
func ConnectDB() {
	host := "localhost"
	port := "27017"
	userName := ""
	password := ""

	connectionStr := ""
	if len(userName) == 0 && len(password) == 0 {
		connectionStr = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		connectionStr = fmt.Sprintf("mongodb://%s:%s@%s:%s", userName, password, host, port)
	}

	dbName := "place-attractions"

	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(connectionStr))
	common.CheckErr(err)
}
