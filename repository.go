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
	Update() bool
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

func (p Pokemon) Update(json Pokemon, id string) bool {
	update := bson.D{{Key: "$set",
	Value: bson.D{
	   {Key: "name", Value: json.Name},
	   {Key: "ability", Value: json.Ability},
	   {Key: "image", Value: json.Image},
	   {Key: "types", Value: json.Types},
	   {Key: "baseStats", Value: json.BaseStats},
   },
	}}
	
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	result, err := initializers.PokemonsCollection.UpdateOne(
		initializers.Context, bson.D{{Key: "_id", Value: objId}}, update)
	if err != nil {
		log.Fatal(err)
	}

	if result.MatchedCount != 0 {
		return true
	}
	return false
}