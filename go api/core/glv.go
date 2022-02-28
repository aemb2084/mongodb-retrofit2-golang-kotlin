package core

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type MongoInfo struct {
	URI        string
	Ctx        context.Context
	Client     *mongo.Client
	Collection *mongo.Collection
}

var MongoDBInfo MongoInfo

// Tipo de variable Pokemones

type Evolution struct {
	Num  string `bson:"num, omitempty" json:"num"`
	Name string `bson:"name, omitempty" json:"name"`
}

type PokemonInfo struct {
	Id            int         `bson:"id,omitempty" json:"id"`
	Num           string      `bson:"num,omitempty" json:"num"`
	Name          string      `bson:"name, omitempty" json:"name"`
	Img           string      `bson:"img, omitempty" json:"img"`
	Type          []string    `bson:"type, omitempty" json:"type"`
	Heigth        string      `bson:"heigth, omitempty" json:"heigth"`
	Weigth        string      `bson:"weigth, omitempty" json:"weigth"`
	Candy         string      `bson:"candy, omitempty" json:"candy"`
	CandyCount    int         `bson:"candy_count, omitempty" json:"candy_count"`
	Egg           string      `bson:"egg, omitempty" json:"egg"`
	SpawnChance   float64     `bson:"spawn_chance, omitempty" json:"spawn_chance"`
	AvgSpawns     float64     `bson:"avg_spawns, omitempty" json:"avg_spawns"`
	SpawnTime     string      `bson:"spawn_time, omitempty" json:"spawn_time"`
	Multipliers   []float64   `bson:"multipliers, omitempty" json:"multipliers"`
	Weaknesses    []string    `bson:"weaknesses, omitempty" json:"weaknesses"`
	PrevEvolution []Evolution `bson:"prev_evolution, omitempty" json:"prev_evolution"`
	NextEvolution []Evolution `bson:"next_evolution, omitempty" json:"next_evolution"`
}

type GeneralHttpResponse struct {
	Status int    `json:"status"`
	Detail string `json:"details"`
}
