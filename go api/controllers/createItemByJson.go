package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"mongo.api/core"
	"mongo.api/utils"
)

func CreateItemByJson(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	var pokemonInfo core.PokemonInfo
	decoded := json.NewDecoder(r.Body)
	err := decoded.Decode(&pokemonInfo)
	utils.CheckError(err)

	result, err := core.MongoDBInfo.Collection.InsertOne(core.MongoDBInfo.Ctx, pokemonInfo)
	utils.CheckError(err)
	httpResponse := core.GeneralHttpResponse{
		200,
		fmt.Sprintf("_id: %v", result.InsertedID),
	}
	json.NewEncoder(w).Encode(httpResponse)

}
