package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongo.api/core"
	"mongo.api/utils"
)

func QueryByIdURL(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	var httpReponse core.GeneralHttpResponse
	collection := core.MongoDBInfo.Collection
	ctx := core.MongoDBInfo.Ctx

	params := mux.Vars(r)
	param := params["_id"]

	_id, err := primitive.ObjectIDFromHex(param)
	utils.CheckError(err)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		httpReponse.Status = 400
		httpReponse.Detail = fmt.Sprintf("_id %v no válido. %v", param, err.Error())
		json.NewEncoder(w).Encode(httpReponse)
		return
	}

	filter := bson.D{{"_id", _id}}

	var pokemonInfo core.PokemonInfo

	err = collection.FindOne(ctx, filter).Decode(&pokemonInfo)
	utils.CheckError(err)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		var Detail struct {
			Status int    `json:"status"`
			Detail string `json:"details"`
		}
		Detail.Status = 404
		Detail.Detail = "No se ha encontrado el _id: " + param + "(" + err.Error() + ")"
		json.NewEncoder(w).Encode(Detail)
		return
	}
	json.NewEncoder(w).Encode(pokemonInfo)

}
