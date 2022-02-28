package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"mongo.api/core"
	"mongo.api/utils"
)

func QueryByIdJson(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	collection := core.MongoDBInfo.Collection
	ctx := core.MongoDBInfo.Ctx

	var httpReponse core.GeneralHttpResponse
	var Id struct {
		Id string `json:"_id"`
	}

	decoded := json.NewDecoder(r.Body)
	err := decoded.Decode(&Id)
	utils.CheckError(err)

	//fmt.Println("_id: ", id)

	_id, err := primitive.ObjectIDFromHex(Id.Id)
	utils.CheckError(err)

	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		httpReponse.Status = 400
		httpReponse.Detail = fmt.Sprintf("_id %v no válido. %v", Id.Id, err.Error())
		json.NewEncoder(w).Encode(httpReponse)	
		return
	}

	filter := bson.D{{"_id", _id}}

	var pokemonInfo core.PokemonInfo
	err = collection.FindOne(ctx, filter).Decode(&pokemonInfo)
	utils.CheckError(err)

	if err != nil{
		w.WriteHeader(http.StatusNotFound)
		httpReponse.Status = 404
		httpReponse.Detail = "No se ha encontrado el _id: " + Id.Id + "(" + err.Error() + ")"
		json.NewEncoder(w).Encode(httpReponse)	
		return
	}

	json.NewEncoder(w).Encode(pokemonInfo)

}
