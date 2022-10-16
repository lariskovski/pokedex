package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	InitialMigration()

	router := gin.Default()
	router.POST("/pokemons", createPokemon)
	router.GET("/pokemons", getPokemon)
	router.GET("/pokemons/:name", getPokemonByName)
	router.DELETE("/pokemons/:name", deletePokemon)
	router.Run("localhost:8080")
}