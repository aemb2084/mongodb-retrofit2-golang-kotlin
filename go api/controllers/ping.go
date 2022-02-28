package controllers

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo/readpref"
	"mongo.api/core"
	"mongo.api/utils"
)

func Ping(w http.ResponseWriter, r *http.Request) {

	err := core.MongoDBInfo.Client.Ping(core.MongoDBInfo.Ctx, readpref.Primary())
	utils.CheckError(err)
	fmt.Fprintf(w, "PING OK")

}
