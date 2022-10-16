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

// Returns all pokemons if no query string requested
func getPokemon(c *gin.Context){
	db, err = gorm.Open("sqlite3", "pokemon.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database.")
	}
	defer db.Close()

	id, ok := c.GetQuery("name")
	if (ok) {
		var pokemon Pokemon
		db.Where("name = ?", id).Find(&pokemon)
		c.IndentedJSON(http.StatusOK, pokemon)
	} else{
		var pokemons []Pokemon
		db.Find(&pokemons)
		c.IndentedJSON(http.StatusOK, pokemons)
	}

}

func deletePokemon(c *gin.Context){
	db, err = gorm.Open("sqlite3", "pokemon.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database.")
	}
	defer db.Close()

	var pokemon Pokemon
	db.Where("name = ?", c.Param("name")).Find(&pokemon)
	db.Delete(&pokemon)

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Pokemon deleted."})
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

func updatePokemon(c *gin.Context) {
	db, err = gorm.Open("sqlite3", "pokemon.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database.")
	}
	defer db.Close()

	var pokemon Pokemon
	var json Pokemon
	
	if err := c.BindJSON(&json); err != nil {
		return
	}

	db.Where("name = ?", c.Param("id")).Find(&pokemon)
	
	db.Model(&pokemon).Select("name", "types", "ability", "image", "baseStats").Updates(Pokemon{
		Name: json.Name,
		Ability: json.Ability,
		Image: json.Image,
		Types: json.Types,
		BaseStats: json.BaseStats,
	})

	c.IndentedJSON(http.StatusAccepted, json.BaseStats)
}
