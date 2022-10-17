package repository

import (
	"errors"

	"github.com/lariskovski/pokedex/api/entity"
)

type PokemonsMemoryDb struct {
	Pokemons []entity.Pokemon
}

type PokemonRepositoryMemory struct {
	db PokemonsMemoryDb
}

func NewPokemonRepositoryMemory(db PokemonsMemoryDb) *PokemonRepositoryMemory{
	return &PokemonRepositoryMemory{db: db}
}

func (p *PokemonRepositoryMemory) Get(name string) (entity.Pokemon, error){
	for _, pokemon := range p.db.Pokemons{
		if pokemon.Name == name {
			return pokemon, nil
		}
	}
	return entity.Pokemon{}, errors.New("Pokemon not found.")
}

func (p *PokemonRepositoryMemory) Create(pokemon entity.Pokemon) (entity.Pokemon, error){
	p.db.Pokemons = append(p.db.Pokemons, pokemon)
	return pokemon, nil
}