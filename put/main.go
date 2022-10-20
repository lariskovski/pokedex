package main

import (
	"context"
	"encoding/json"
	"fmt"
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
	lambda.Start(postPokemon)
}


// Returns all pokemons if no query string requested
func postPokemon(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.PathParameters["name"]
	if name != "" {
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
		update := bson.D{{Key: "$set",
			Value: bson.D{
				{Key: "ability", Value: pokemon.Ability},
				{Key: "image", Value: pokemon.Image},
				{Key: "types", Value: pokemon.Types},
				{Key: "baseStats", Value: pokemon.BaseStats},
			},
		}}
		result, err := PokemonsCollection.UpdateOne(Context, bson.D{{Key: "name", Value: name}} , update)
		if err != nil {
			log.Fatal(err)
		}

		if result.MatchedCount != 0 {
			// Transform Pokemon struct into json
			pokemon.Name = name
			pokemon.Id = fmt.Sprint(result.UpsertedID)
			json, err := json.Marshal(pokemon)
			if err != nil {
				log.Fatal(err)
			}
			return events.APIGatewayProxyResponse{StatusCode: 201, Body: string(json)}, nil
		} else {
			json, err := json.Marshal(Response{Message: "No match found."})
			if err != nil {
				log.Fatal(err)
			}
			return events.APIGatewayProxyResponse{StatusCode: 404, Body: string(json)}, nil
		}
	} else {
		json, err := json.Marshal(Response{Message: "Missing name path parameter."})
		if err != nil {
			log.Fatal(err)
		}
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: string(json)}, nil
	}
}

