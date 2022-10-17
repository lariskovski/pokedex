package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lariskovski/pokedex/api/entity"
	"github.com/lariskovski/pokedex/api/repository"
	"github.com/lariskovski/pokedex/api/service"
)

func CreatePokemon(c *gin.Context) {
	db := repository.PokemonsMemoryDb{Pokemons: []entity.Pokemon{}}
	repositoryMemory := repository.NewPokemonRepositoryMemory(db)
	service := service.NewPokemonService(repositoryMemory)

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

