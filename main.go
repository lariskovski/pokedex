package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lariskovski/pokedex/api/controller"
	"github.com/lariskovski/pokedex/api/initializers"
)

func init(){
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main(){
	router := gin.Default()
	router.POST("/pokemons", controller.CreatePokemon)
	router.GET("/pokemons", controller.GetPokemon)
	// router.PUT("/pokemons/:id", updatePokemon)
	// router.DELETE("/pokemons/:id", deletePokemon)
	router.Run("localhost:8080")
}