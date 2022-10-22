package commons

import (
	"context"
	"log"
	"os"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Collection *mongo.Collection
	Context context.Context
)

func Connect(){
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
	if err != nil {
		log.Fatal(err)
	}
	Context = context.Background()
	err = client.Connect(Context)
	if err != nil {
		log.Fatal(err)
	}
	pokedexDB := client.Database("pokedex")
	Collection = pokedexDB.Collection("pokemon")
}