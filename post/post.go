package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/mongo"
	uuid "github.com/satori/go.uuid"
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


func main() {
	lambda.Start(postPokemon)
}


// Returns all pokemons if no query string requested
func postPokemon(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
	
	
	var pokemon Pokemon
	json.Unmarshal([]byte(request.Body), &pokemon)
	pokemon.Id = uuid.NewV4().String()

	_, err = PokemonsCollection.InsertOne(Context, pokemon)
	if err != nil {
		log.Fatal(err)
	}
	return events.APIGatewayProxyResponse{StatusCode: 201}, nil
}

