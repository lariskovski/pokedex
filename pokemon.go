package main

import (
	// "fmt"
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



type Pokemon struct {
	Name string `json:"name"`
	Types []string `json:"types"`
	Image string `json:"image"`
	Ability string `json:"ability"`
	BaseStats map[string]string `json:"baseStats"`
}


// Returns all pokemons if no query string requested
func getPokemon(c *gin.Context){
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	pokedexDB := client.Database("pokedex")
	PokemonsCollection := pokedexDB.Collection("pokemon")
	
	// If query string name is present returns one value only
	// or all values
	name, ok := c.GetQuery("name")
	if (ok) {
		cursor, err := PokemonsCollection.Find(ctx, bson.M{"name": name})
		if err != nil {
			log.Fatal(err)
		}
		var pokemon []bson.M
		if err = cursor.All(ctx, &pokemon); err != nil {
			log.Fatal(err)
		}
		c.IndentedJSON(http.StatusOK, pokemon)

	} else {
		cursor, err := PokemonsCollection.Find(ctx, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		var pokemons []bson.M
		if err = cursor.All(ctx, &pokemons); err != nil {
			log.Fatal(err)
		}
		c.IndentedJSON(http.StatusOK, pokemons)
	}
}

func deletePokemon(c *gin.Context){
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	pokedexDB := client.Database("pokedex")
	PokemonsCollection := pokedexDB.Collection("pokemon")
	
	deleteResult, err := PokemonsCollection.DeleteOne(ctx, bson.M{"name": c.Param("name")})
	if err != nil {
		log.Fatal(err)
	}

	if deleteResult != nil {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Pokemon deleted."})
	}
}

func createPokemon(c *gin.Context) {
	var pokemon Pokemon

	if err := c.BindJSON(&pokemon); err != nil {
		return
	}
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	pokedexDB := client.Database("pokedex")
	// err = pokedexDB.CreateCollection(ctx, "pokemon")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	PokemonsCollection := pokedexDB.Collection("pokemon")
	// defer pokemonsCollection.Drop(ctx)
	result, err := PokemonsCollection.InsertOne(ctx, pokemon)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusCreated, result)
}

func updatePokemon(c *gin.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	pokedexDB := client.Database("pokedex")
	PokemonsCollection := pokedexDB.Collection("pokemon")

	var json Pokemon
	if err := c.BindJSON(&json); err != nil {
		return
	}
	update := bson.D{{Key: "$set",
		 Value: bson.D{
			{Key: "name", Value: json.Name},
			{Key: "ability", Value: json.Ability},
			{Key: "image", Value: json.Image},
			{Key: "types", Value: json.Types},
			{Key: "baseStats", Value: json.BaseStats},
		},
	}}
	result, err := PokemonsCollection.UpdateOne(ctx, bson.D{{Key: "name", Value: c.Param("name")}} , update)
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount != 0 {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Pokemon updated."})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No match found."})
	}

}
