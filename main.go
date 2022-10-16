package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoURI = os.Getenv("MONGODB_URI")
	PokemonsCollection *mongo.Collection
	ctx context.Context
)

func main(){

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx = context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	pokedexDB := client.Database("pokedex")
	PokemonsCollection = pokedexDB.Collection("pokemon")

	router := gin.Default()
	router.POST("/pokemons", createPokemon)
	router.GET("/pokemons", getPokemon)
	router.PUT("/pokemons/:id", updatePokemon)
	router.DELETE("/pokemons/:id", deletePokemon)
	router.Run("localhost:8080")
}