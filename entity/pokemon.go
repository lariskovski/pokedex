package entity

import (
	uuid "github.com/satori/go.uuid"
)

type PokemonRepository interface {
	Get(name string) (Pokemon, error)
	Create(pokemon Pokemon) (Pokemon, error)
	Update(id string, pokemon Pokemon) (Pokemon, error)
	Delete(id string) error
}

type Pokemon struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Types []string `json:"types"`
	Image string `json:"image"`
	Ability string `json:"ability"`
	BaseStats map[string]string `json:"baseStats"`
}

func NewPokemon() *Pokemon {
	pokemon := Pokemon {
		Id: uuid.NewV4().String(),
	}
	return &pokemon
}
