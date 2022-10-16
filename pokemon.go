package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gorm.io/datatypes"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	
)

var db *gorm.DB
var err error


type Pokemon struct {
	gorm.Model
	Name string `json:"name"`
	Types datatypes.JSON `json:"types"`
	Image string `json:"image"`
	Ability string `json:"ability"`
	BaseStats datatypes.JSON `json:"baseStats"`
}

func InitialMigration() {
	db, err = gorm.Open("sqlite3", "pokemon.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database.")
	}
	defer db.Close()
	
	db.AutoMigrate(&Pokemon{})
	fmt.Println("Successfully created database.")
}

func getPokemon(c *gin.Context){
	db, err = gorm.Open("sqlite3", "pokemon.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database.")
	}
	defer db.Close()

	var pokemons []Pokemon
	db.Find(&pokemons)

	c.IndentedJSON(http.StatusOK, pokemons)
}

func getPokemonByName(c *gin.Context){
	db, err = gorm.Open("sqlite3", "pokemon.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database.")
	}
	defer db.Close()

	var pokemon Pokemon
	db.Where("name = ?", c.Param("name")).Find(&pokemon)
	if (pokemon.ID != 0) {
		c.IndentedJSON(http.StatusOK, pokemon)
	} else{
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Pokemon not found."})
	}
}

func createPokemon(c *gin.Context) {
	var pokemon Pokemon

	if err := c.BindJSON(&pokemon); err != nil {
		return
	}
	db, err = gorm.Open("sqlite3", "pokemon.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database.")
	}
	defer db.Close()

	c.IndentedJSON(http.StatusCreated, db.Create(&Pokemon{
		Name: pokemon.Name,
		Ability: pokemon.Ability,
		Image: pokemon.Image,
		Types: pokemon.Types,
		BaseStats: pokemon.BaseStats,
	}))
}

