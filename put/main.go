package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Pokemon struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Types []string `json:"types"`
	Image string `json:"image"`
	Ability string `json:"ability"`
	BaseStats map[string]string `json:"baseStats"`
}

type Response struct {
	Message string `json:"message"`
}

func main() {
	lambda.Start(putPokemon)
}


func putPokemon(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Name is not null if there is the name path on the API
	// PUT /pokemons/{name}
	name := request.PathParameters["name"]
	if name != "" {
		// MongoDB config
		client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_URI")))
		if err != nil {
			log.Fatal(err)
		}
		Context := context.Background()
		err = client.Connect(Context)
		if err != nil {
			log.Fatal(err)
		}
		PokemonsCollection := client.Database("pokedex").Collection("pokemon")
		
		// Transforms request body json into Pokemon struct
		var pokemon Pokemon
		json.Unmarshal([]byte(request.Body), &pokemon)
		// Creates the update object to pass into Mongo
		update := bson.D{{Key: "$set",
			Value: bson.D{
				{Key: "ability", Value: pokemon.Ability},
				{Key: "image", Value: pokemon.Image},
				{Key: "types", Value: pokemon.Types},
				{Key: "baseStats", Value: pokemon.BaseStats},
			},
		}}

		// Updates object and returns the new one into pokemon
		err = PokemonsCollection.FindOneAndUpdate(Context, bson.D{{Key: "name", Value: name}}, update).Decode(&pokemon)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				// This error means your query did not match any documents.
				json, err := json.Marshal(Response{Message: "No match found."})
				if err != nil {
					log.Fatal(err)
				}
				return events.APIGatewayProxyResponse{StatusCode: 404, Body: string(json)}, nil
			}
		}
		// No error on FindOneAndUpdate returns 200
		json, err := json.Marshal(pokemon)
		if err != nil {
			log.Fatal(err)
		}
		return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(json)}, nil
	} else {
		// Enter when no {name} is passed into API call
		json, err := json.Marshal(Response{Message: "Missing name path parameter."})
		if err != nil {
			log.Fatal(err)
		}
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: string(json)}, nil
	}
}

