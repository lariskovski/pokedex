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
	router.Run("localhost:8080")
}