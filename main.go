package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	// "errors"
)

type pokemon struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Types []string `json:"types"`
	Image string `json:"image"`
	Ability string `json:"ability"`
	BaseStats map[string]string `json:"baseStats"`
}

var pokemons = []pokemon{
	{Id: "#001", Name: "Bulbasaur", Image: "/swordshield/pokemon/small/001.png", Types: []string{"grass", "poison"}, Ability: "Overgrow Chlorophyll", BaseStats: map[string]string{"healthPoints": "45", "attack": "49", "defense": "49", "speedAttack": "65", "speedDefense": "65", "specialDefense": "45"}},
	{Id: "#002", Name: "Ivysaur", Image: "/swordshield/pokemon/small/002.png", Types: []string{"grass", "poison"}, Ability: "Overgrow Chlorophyll", BaseStats: map[string]string{"healthPoints": "60", "attack": "62", "defense": "63", "speedAttack": "80", "speedDefense": "80", "specialDefense": "60"}},
}

func getPokemon(c *gin.Context){
	c.IndentedJSON(http.StatusOK, pokemons)

}

func main(){
	router := gin.Default()
	router.GET("pokemons", getPokemon)
	router.Run("localhost:8080")
}