package entity

type PokemonRepository interface {
	Get(name string) (Pokemon, error)
	Create(pokemon Pokemon) (Pokemon, error)
}

type Pokemon struct {
	Name string `json:"name"`
	Types []string `json:"types"`
	Image string `json:"image"`
	Ability string `json:"ability"`
	BaseStats map[string]string `json:"baseStats"`
}

func NewPokemon() *Pokemon {
	pokemon := Pokemon {}
	return &pokemon
}
