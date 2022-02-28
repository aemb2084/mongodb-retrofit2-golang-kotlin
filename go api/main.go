package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mongo.api/core"
	"mongo.api/routers"
	"mongo.api/utils"
)

func main() {

	core.MongoDBInfo.URI = "mongodb+srv://admin:administrator@cluster0.gjj7f.mongodb.net/test"

	//ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	core.MongoDBInfo.Ctx = ctx

	client, err := mongo.Connect(core.MongoDBInfo.Ctx, options.Client().ApplyURI(core.MongoDBInfo.URI))
	utils.CheckError(err)
	defer client.Disconnect(core.MongoDBInfo.Ctx)

	core.MongoDBInfo.Client = client
	core.MongoDBInfo.Collection = core.MongoDBInfo.Client.Database("mongodb").Collection("pokemon")

	test := fmt.Sprint(core.MongoDBInfo.Client)

	var myClient uuid.UUID
	json.Unmarshal([]byte(test), &myClient)
	fmt.Println("Mongo client: ", myClient)

	fmt.Println("Server is running at :8080/")
	server := http.ListenAndServe(":8080", routers.NewRoute())
	log.Fatal(server)

}
