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
	Name string `json:"name"`
	Types []string `json:"types"`
	Image string `json:"image"`
	Ability string `json:"ability"`
	BaseStats map[string]string `json:"baseStats"`
}


func main() {
	lambda.Start(getPokemon)
}


// Returns all pokemons if no query string requested
func getPokemon(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	mongoURI := os.Getenv("MONGODB_URI")
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}
	Context := context.Background()
	err = client.Connect(Context)
	if err != nil {
		log.Fatal(err)
	}
	PokemonsCollection := client.Database("pokedex").Collection("pokemon")
	
	
	// If query string name is present returns one value only
	// or all values
	name := request.QueryStringParameters["name"]
	if (name != "") {
		var pokemon Pokemon
		err := PokemonsCollection.FindOne(Context, bson.M{"name": name}).Decode(&pokemon)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				// This error means your query did not match any documents.
				return events.APIGatewayProxyResponse{StatusCode: 404}, nil
			}
		}
		// Transform Mongo response into json
		json, err := json.Marshal(pokemon)
		if err != nil {
			log.Fatal(err)
		}
		return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(json)}, nil
	} else {
		result, err := PokemonsCollection.Find(Context, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		var pokemons []Pokemon
		if err = result.All(Context, &pokemons); err != nil {
			log.Fatal(err)
		}
		// Transform mongo response into json
		body, err := json.Marshal(pokemons)
		if err != nil {
			log.Fatal(err)
		}
		return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(body)}, nil
	}
}

