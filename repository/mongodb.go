package repository

import (
	"context"
	"errors"
	"log"

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


func (p *PokemonRepositoryMongoDb) Update(id string, pokemon entity.Pokemon) (entity.Pokemon, error){
	update := bson.D{{Key: "$set",
	Value: bson.D{
	   {Key: "name", Value: pokemon.Name},
	   {Key: "ability", Value: pokemon.Ability},
	   {Key: "image", Value: pokemon.Image},
	   {Key: "types", Value: pokemon.Types},
	   {Key: "baseStats", Value: pokemon.BaseStats},
   },
	}}
	
	result, err := p.db.UpdateOne(
		p.ctx, bson.D{{Key: "id", Value: id}}, update)
	if err != nil {
		log.Fatal(err)
	}
	if result.MatchedCount != 0 {
		return pokemon, nil
	}
	return pokemon, errors.New("No match found for update.")
}

func (p *PokemonRepositoryMongoDb) Delete(id string) error {
	deleteResult, err := p.db.DeleteOne(
		p.ctx, bson.D{{Key: "id", Value: id }})
	if deleteResult.DeletedCount != 0 {
		return nil
	} else {
		return err
	}
}