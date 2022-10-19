package main

import (
	"os"
	"log"
	"context"
	"encoding/json"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
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
		result, err := PokemonsCollection.Find(Context, bson.M{"name": name})
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

