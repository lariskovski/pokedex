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

type Response struct {
	Message string `json:"message"`
}

func main() {
	lambda.Start(postPokemon)
}


// Returns all pokemons if no query string requested
func postPokemon(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := request.Body
	if !(body == "" || body == "{}") {
		// Mongo cofig
		client, err := mongo.NewClient(options.Client().ApplyURI( os.Getenv("MONGODB_URI")))
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
		pokemon.Id = uuid.NewV4().String()
	
		_, err = PokemonsCollection.InsertOne(Context, pokemon)
		if err != nil {
			log.Fatal(err)
		}
	
		// Transform Pokemon struct into json
		json, err := json.Marshal(pokemon)
		if err != nil {
			log.Fatal(err)
		}
		return events.APIGatewayProxyResponse{StatusCode: 201, Body: string(json)}, nil
	} else {
		json, err := json.Marshal(Response{Message: "Body is empty."})
		if err != nil {
			log.Fatal(err)
		}
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: string(json)}, nil
	}
}

