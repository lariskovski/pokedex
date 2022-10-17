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
func (p *PokemonRepositoryMemory) Update(pokemon entity.Pokemon) (entity.Pokemon, error){
	for index, poke := range p.db.Pokemons{
		if poke.Name == pokemon.Name {
			p.db.Pokemons[index] = pokemon
			return pokemon, nil
		}
	}
	return entity.Pokemon{}, errors.New("Pokemon not found.")
}
func (p *PokemonRepositoryMemory) Delete(pokemon entity.Pokemon) bool {
	for index, poke := range p.db.Pokemons{
		if poke.Name == pokemon.Name {
			p.db.Pokemons[index] = entity.Pokemon{}
			return true
		}
	}
	return false
}