package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/mgo.v2/bson"
	"mongo.api/core"
	"mongo.api/utils"
)

func DeleteOneById(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	var httpResponse core.GeneralHttpResponse

	collection := core.MongoDBInfo.Collection
	ctx := core.MongoDBInfo.Ctx

	params := mux.Vars(r)
	param := params["_id"]

	_id, err := primitive.ObjectIDFromHex(param)
	utils.CheckError(err)
	if err != nil {
		httpResponse.Status = 400
		httpResponse.Detail = fmt.Sprintf("_id %v no válido. %v", _id, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(httpResponse)
		return
	}

	//filter := bson.D{{"_id", _id}}

	result, err := collection.DeleteOne(ctx, bson.M{"_id": _id})
	utils.CheckError(err)

	if result.DeletedCount == 0 {
		httpResponse.Status = 404
		httpResponse.Detail = "No se ha eliminado ningún item."
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(httpResponse)
		return
	}

	httpResponse.Status = 200
	httpResponse.Detail = fmt.Sprintf("Item con _id: %v eliminado exitosamente.", param)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(httpResponse)

}
