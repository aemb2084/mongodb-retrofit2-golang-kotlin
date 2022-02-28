package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"mongo.api/core"
	"mongo.api/utils"
)

func CreateClient(contextBck context.Context) core.MongoInfo {

	var clientData core.MongoInfo

	clientData.URI = "mongodb+srv://admin:administrator@cluster0.gjj7f.mongodb.net/test"

	ctx, cancel := context.WithTimeout(contextBck, 10*time.Second)
	defer cancel()
	clientData.Ctx = ctx

	client, err := mongo.Connect(clientData.Ctx, options.Client().ApplyURI(clientData.URI))
	utils.CheckError(err)
	defer client.Disconnect(ctx)

	clientData.Client = client

	return clientData

}
