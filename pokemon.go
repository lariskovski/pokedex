package main

import (
	// "fmt"
	"net/http"
	"context"
	"log"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/gin-gonic/gin"
)



type Pokemon struct {
	Name string `json:"name"`
	Types []string `json:"types"`
	Image string `json:"image"`
	Ability string `json:"ability"`
	BaseStats map[string]string `json:"baseStats"`
}


// // Returns all pokemons if no query string requested
// func getPokemon(c *gin.Context){
// 	db, err = gorm.Open("sqlite3", "pokemon.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("Failed to connect to database.")
// 	}
// 	defer db.Close()

// 	id, ok := c.GetQuery("name")
// 	if (ok) {
// 		var pokemon Pokemon
// 		db.Where("name = ?", id).Find(&pokemon)
// 		c.IndentedJSON(http.StatusOK, pokemon)
// 	} else{
// 		var pokemons []Pokemon
// 		db.Find(&pokemons)
// 		c.IndentedJSON(http.StatusOK, pokemons)
// 	}

// }

// func deletePokemon(c *gin.Context){
// 	db, err = gorm.Open("sqlite3", "pokemon.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("Failed to connect to database.")
// 	}
// 	defer db.Close()

// 	var pokemon Pokemon
// 	db.Where("name = ?", c.Param("name")).Find(&pokemon)
// 	db.Delete(&pokemon)

// 	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Pokemon deleted."})
// }

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

// func updatePokemon(c *gin.Context) {
// 	db, err = gorm.Open("sqlite3", "pokemon.db")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		panic("Failed to connect to database.")
// 	}
// 	defer db.Close()

// 	var pokemon Pokemon
// 	var json Pokemon
	
// 	if err := c.BindJSON(&json); err != nil {
// 		return
// 	}

// 	db.Where("name = ?", c.Param("id")).Find(&pokemon)
	
// 	db.Model(&pokemon).Select("name", "types", "ability", "image", "baseStats").Updates(Pokemon{
// 		Name: json.Name,
// 		Ability: json.Ability,
// 		Image: json.Image,
// 		Types: json.Types,
// 		BaseStats: json.BaseStats,
// 	})

// 	c.IndentedJSON(http.StatusAccepted, json.BaseStats)
// }
