package main

import (
	"fmt"

	"github.com/lariskovski/pokedex/api/entity"
	"github.com/lariskovski/pokedex/api/repository"
	"github.com/lariskovski/pokedex/api/service"
)

// "github.com/gin-gonic/gin"
// "github.com/lariskovski/pokedex/api/initializers"

// func init(){
// 	initializers.LoadEnvVars()
// 	initializers.ConnectToDB()
// }

func main(){
	db := repository.PokemonsMemoryDb{Pokemons: []entity.Pokemon{}}
	repositoryMemory := repository.NewPokemonRepositoryMemory(db)

	service := service.NewPokemonService(repositoryMemory)

	pok, _ := service.Create("Pikachu", "Shock Wave", []string{"Electric"}, "pikachu.png", map[string]string{"hp": "40"})
	fmt.Print(pok)
	// router := gin.Default()
	// router.POST("/pokemons", createPokemon)
	// router.GET("/pokemons", getPokemon)
	// router.PUT("/pokemons/:id", updatePokemon)
	// router.DELETE("/pokemons/:id", deletePokemon)
	// router.Run("localhost:8080")
}