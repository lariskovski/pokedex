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
	router.PUT("/pokemons/:id", updatePokemon)
	router.DELETE("/pokemons/:id", deletePokemon)
	router.Run("localhost:8080")
}