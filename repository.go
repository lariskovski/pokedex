package main

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"github.com/lariskovski/pokedex/api/initializers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	Create() *mongo.InsertOneResult
    GetAll() []primitive.M
    GetByName() []bson.M
	// Update() *mongo.UpdateResult
	// Delete() *mongo.DeleteResult
}

func (p Pokemon) Create() *mongo.InsertOneResult {
	result, err := initializers.PokemonsCollection.InsertOne(initializers.Context, p)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func (p Pokemon) GetAll() []bson.M {
	cursor, err := initializers.PokemonsCollection.Find(initializers.Context, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var pokemons []bson.M
	if err = cursor.All(initializers.Context, &pokemons); err != nil {
		log.Fatal(err)
	}
	return pokemons
}

func (p Pokemon) GetByName(name string) []bson.M {
	cursor, err := initializers.PokemonsCollection.Find(initializers.Context, bson.M{"name": name})
	if err != nil {
		log.Fatal(err)
	}
	var pokemon []bson.M
	if err = cursor.All(initializers.Context, &pokemon); err != nil {
		log.Fatal(err)
	}
	return pokemon
}