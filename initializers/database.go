package initializers

import (
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var (
	PokemonsCollection *mongo.Collection
	Context context.Context
)

func ConnectToDB(){
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	Context = context.Background()
	err = client.Connect(Context)
	if err != nil {
		log.Fatal(err)
	}
	pokedexDB := client.Database("pokedex")
	PokemonsCollection = pokedexDB.Collection("pokemon")
}