package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	db "github.com/lariskovski/pokedex/api/commons"
)

type Pokemon struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Types []string `json:"types"`
	Image string `json:"image"`
	Ability string `json:"ability"`
	BaseStats map[string]string `json:"baseStats"`
}

func init(){
	db.Connect()
}

func main() {
	lambda.Start(getPokemon)
}

// Returns all pokemons if no query string requested
func getPokemon(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If query string name is present returns one value only
	// or all values
	name := request.QueryStringParameters["name"]
	if (name != "") {
		var pokemon Pokemon
		err := db.Collection.FindOne(db.Context, bson.M{"name": name}).Decode(&pokemon)
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
		result, err := db.Collection.Find(db.Context, bson.M{})
		if err != nil {
			log.Fatal(err)
		}
		var pokemons []Pokemon
		if err = result.All(db.Context, &pokemons); err != nil {
			log.Fatal(err)
		}
		// Transform mongo response into json
		json, err := json.Marshal(pokemons)
		if err != nil {
			log.Fatal(err)
		}
		return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(json)}, nil
	}
}

