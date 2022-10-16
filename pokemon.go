package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lariskovski/pokedex/api/initializers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Pokemon struct {
	Name string `json:"name"`
	Types []string `json:"types"`
	Image string `json:"image"`
	Ability string `json:"ability"`
	BaseStats map[string]string `json:"baseStats"`
}

func createPokemon(c *gin.Context) {
	var pokemon Pokemon
	if err := c.BindJSON(&pokemon); err != nil {
		return
	}
	result := pokemon.Create()
	c.IndentedJSON(http.StatusCreated, result)
}


// Returns all pokemons if no query string requested
func getPokemon(c *gin.Context){
	// If query string name is present returns one value only
	// or all values
	name, ok := c.GetQuery("name")
	var pokemon Pokemon
	if (ok) {
		result := pokemon.GetByName(name)
		c.IndentedJSON(http.StatusOK, result)

	} else {
		result := pokemon.GetAll()
		c.IndentedJSON(http.StatusOK, result)
	}
}


func updatePokemon(c *gin.Context) {
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
		
	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	result, err := initializers.PokemonsCollection.UpdateOne(initializers.Context, bson.D{{Key: "_id", Value: objId}} , update)
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount != 0 {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Pokemon updated."})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No match found."})
	}
}


func deletePokemon(c *gin.Context){
	objId, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		log.Fatal(err)
	}
	deleteResult, err := initializers.PokemonsCollection.DeleteOne(initializers.Context, bson.D{{Key: "_id", Value: objId }})
	if err != nil {
		log.Fatal(err)
	}

	if deleteResult.DeletedCount != 0 {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Pokemon deleted."})
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No match found."})
	}
}
