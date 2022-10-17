package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lariskovski/pokedex/api/entity"
	"github.com/lariskovski/pokedex/api/initializers"
)

func CreatePokemon(c *gin.Context) {
	// db := repository.PokemonsMemoryDb{Pokemons: []entity.Pokemon{}}
	// repositoryMemory := repository.NewPokemonRepositoryMemory(db)
	// service := initializers.Service.NewPokemonService(repositoryMemory)

	var pokemon entity.Pokemon
	if err := c.BindJSON(&pokemon); err != nil {
		return
	}
	
	result, err := initializers.Service.Create(
		pokemon.Name,
		pokemon.Ability,
		pokemon.Types,
		pokemon.Image,
		pokemon.BaseStats,
	)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusCreated, result)
}

func GetPokemon(c *gin.Context){
	name, ok := c.GetQuery("name")
	if (ok) {
		result, err := initializers.Service.FindByName(name)
		if err != nil{
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No match found."})
			return
		}
		c.IndentedJSON(http.StatusOK, result)
	}
}


func DeletePokemon(c *gin.Context){
	if err := initializers.Service.Delete(c.Param("id")); err == nil {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Accepted"})
	} else {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error"})
	}
}


func UpdatePokemon(c *gin.Context) {
	var pokemon entity.Pokemon
	if err := c.BindJSON(&pokemon); err != nil {
		return
	}
	_, err := initializers.Service.Update(c.Param("id"), pokemon.Name,
	pokemon.Ability,
	pokemon.Types,
	pokemon.Image,
	pokemon.BaseStats,)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No match found."})
	} else {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Pokemon updated."})
	}
}