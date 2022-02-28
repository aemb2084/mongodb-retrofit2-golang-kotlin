package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"mongo.api/core"
	"mongo.api/utils"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "application/json")

	fmt.Println("createItemByStructure")

	var pokemonInfo core.PokemonInfo
	pokemonInfo.AvgSpawns = 0.25
	pokemonInfo.Candy = "algun dulce"
	pokemonInfo.CandyCount = 20
	pokemonInfo.Egg = "Algun huevo"
	pokemonInfo.Heigth = "1.76 m"
	pokemonInfo.Id = 123456
	pokemonInfo.Img = "www.google.com"
	pokemonInfo.Multipliers = []float64{0.12, 2.0, 3.25, 0.25}
	pokemonInfo.Name = "super picachu"
	pokemonInfo.PrevEvolution = []core.Evolution{
		{"001", "Super picachu 1"},
		{"002", "Super picachu 2"},
	}
	pokemonInfo.NextEvolution = []core.Evolution{
		{"003", "Super picachu 3"},
		{"004", "Super picachu 4"},
	}
	pokemonInfo.Num = "151"
	pokemonInfo.SpawnChance = 1.25
	pokemonInfo.SpawnTime = "1 hora"
	pokemonInfo.Type = []string{
		"Atleta",
		"Corredor",
		"Saltarin",
	}

	coll := core.MongoDBInfo.Collection
	ctx := core.MongoDBInfo.Ctx

	result, err := coll.InsertOne(ctx, pokemonInfo)
	fmt.Println("Result: ", result.InsertedID)
	utils.CheckError(err)
	httpResponse := core.GeneralHttpResponse{
		200,
		fmt.Sprintf("_id: %v", result.InsertedID),
	} 
	json.NewEncoder(w).Encode(httpResponse)

}
