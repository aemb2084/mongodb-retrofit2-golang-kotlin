package controllers

import (
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2/bson"
	"mongo.api/core"
	"mongo.api/utils"
)

func Dblist(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")
	databases, err := core.MongoDBInfo.Client.ListDatabaseNames(core.MongoDBInfo.Ctx, bson.M{})
	utils.CheckError(err)
	json.NewEncoder(w).Encode(databases)

}
