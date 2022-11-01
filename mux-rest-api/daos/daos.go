package daos

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var url = "mongodb://localhost:27017"
var DBNAME = "ems"

func initClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Client intialize ...")
	}
	return client
}

type Entity interface {
	GetCollectionName() string
}

func DoSome(f func(collection *mongo.Collection, ctx context.Context) error, e Entity) error {
	client := initClient()
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	collection := client.Database(DBNAME).Collection(e.GetCollectionName())
	return f(collection, ctx)
}
