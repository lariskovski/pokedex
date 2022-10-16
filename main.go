package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

var (
	mongoURI = os.Getenv("MONGODB_URI")
)

func main(){

	router := gin.Default()
	router.POST("/pokemons", createPokemon)
	router.GET("/pokemons", getPokemon)
	router.PUT("/pokemons/:name", updatePokemon)
	router.DELETE("/pokemons/:name", deletePokemon)
	router.Run("localhost:8080")


}