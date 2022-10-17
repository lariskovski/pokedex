package repository

import (
	"context"
	"github.com/lariskovski/pokedex/api/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PokemonRepositoryMongoDb struct {
	db *mongo.Collection
	ctx context.Context
}

func NewPokemonRepositoryMongoDb(db *mongo.Collection) *PokemonRepositoryMongoDb{
	return &PokemonRepositoryMongoDb{db: db}
}

func (p *PokemonRepositoryMongoDb) Create(pokemon entity.Pokemon) (entity.Pokemon, error){
	_, err := p.db.InsertOne(p.ctx, pokemon)
	return pokemon, err
}

func (p *PokemonRepositoryMongoDb) Get(name string) (entity.Pokemon, error){
	var pokemon entity.Pokemon
	cursor := p.db.FindOne(p.ctx, bson.M{"name": name})
	err := cursor.Decode(&pokemon)
	return pokemon, err
}


// func (p *PokemonRepositoryMongoDb) Update(pokemon entity.Pokemon) (entity.Pokemon, error){

// }

func (p *PokemonRepositoryMongoDb) Delete(id string) error {
	deleteResult, err := p.db.DeleteOne(
		p.ctx, bson.D{{Key: "id", Value: id }})
	if deleteResult.DeletedCount != 0 {
		return nil
	} else {
		return err
	}
}