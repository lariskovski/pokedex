package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	InitialMigration()

	router := gin.Default()
	router.POST("/pokemons", createPokemon)
	router.GET("/pokemons", getPokemon)
	router.DELETE("/pokemons/:name", deletePokemon)
	router.PUT("/pokemons/:id", updatePokemon)
	router.Run("localhost:8080")
}