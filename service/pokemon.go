package service

import (
	"github.com/lariskovski/pokedex/api/entity"
)

type PokemonService struct {
	Repository entity.PokemonRepository
}

func NewPokemonService(repository entity.PokemonRepository) *PokemonService {
	return &PokemonService{Repository: repository}
}

func (p *PokemonService) FindByName(name string) (entity.Pokemon, error){
	pokemon, err := p.Repository.Get(name)
	if err != nil {
		return pokemon, err
	}
	return pokemon, nil
}

func (p *PokemonService) Create(name string, ability string, types []string, image string, baseStats map[string]string ) (entity.Pokemon, error){
	pokemon := entity.NewPokemon()
	pokemon.Name = name
	pokemon.Ability = ability
	pokemon.Image = image
	pokemon.Types = types
	pokemon.BaseStats = baseStats

	result, err := p.Repository.Create(*pokemon)
	return result, err
}