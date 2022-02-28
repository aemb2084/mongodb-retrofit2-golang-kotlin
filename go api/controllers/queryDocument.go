package controllers

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"mongo.api/core"
	"mongo.api/utils"
)

func QueryDocument(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	var pokemonInfo []core.PokemonInfo

	collection := core.MongoDBInfo.Collection
	ctx := core.MongoDBInfo.Ctx

	cur, err := collection.Find(ctx, bson.D{})
	utils.CheckError(err)
	defer cur.Close(ctx)

	for cur.Next(ctx) {
		var temp core.PokemonInfo
		err := cur.Decode(&temp)
		pokemonInfo = append(pokemonInfo, temp)
		utils.CheckError(err)
	}

	json.NewEncoder(w).Encode(pokemonInfo)

}
