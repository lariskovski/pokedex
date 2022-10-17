package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lariskovski/pokedex/api/entity"
	"github.com/lariskovski/pokedex/api/initializers"
	"github.com/lariskovski/pokedex/api/repository"
	"github.com/lariskovski/pokedex/api/service"
)

func CreatePokemon(c *gin.Context) {
	// db := repository.PokemonsMemoryDb{Pokemons: []entity.Pokemon{}}
	// repositoryMemory := repository.NewPokemonRepositoryMemory(db)
	// service := service.NewPokemonService(repositoryMemory)

	db := initializers.PokemonsCollection
	repository := repository.NewPokemonRepositoryMongoDb(db)
	service := service.NewPokemonService(repository)
	var pokemon entity.Pokemon

	if err := c.BindJSON(&pokemon); err != nil {
		return
	}
	result, err := service.Create(
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
	db := initializers.PokemonsCollection
	repository := repository.NewPokemonRepositoryMongoDb(db)
	service := service.NewPokemonService(repository)

	name, ok := c.GetQuery("name")
	if (ok) {
		result, err := service.FindByName(name)
		if err != nil{
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No match found."})
			return
		}
		c.IndentedJSON(http.StatusOK, result)
	}
}


func DeletePokemon(c *gin.Context){
	db := initializers.PokemonsCollection
	repository := repository.NewPokemonRepositoryMongoDb(db)
	service := service.NewPokemonService(repository)

	if err := service.Delete(c.Param("id")); err == nil {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Accepted"})
	} else {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Error"})
	}
}


func UpdatePokemon(c *gin.Context) {
	db := initializers.PokemonsCollection
	repository := repository.NewPokemonRepositoryMongoDb(db)
	service := service.NewPokemonService(repository)

	var pokemon entity.Pokemon
	if err := c.BindJSON(&pokemon); err != nil {
		return
	}
	_, err := service.Update(c.Param("id"), pokemon.Name,
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