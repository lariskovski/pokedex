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

type Response struct {
	Message string `json:"message"`
}

func main() {
	lambda.Start(deletePokemon)
}


func deletePokemon(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
	
	name := request.PathParameters["name"]

	if name != "" {
		deleteResult, err := PokemonsCollection.DeleteOne(Context, bson.D{{Key: "name", Value: name }})
		if err != nil {
			log.Fatal(err)
		}
	
		if deleteResult.DeletedCount != 0 {
			json, err := json.Marshal(Response{Message: "Pokemon deleted."})
			if err != nil {
				log.Fatal(err)
			}
			return events.APIGatewayProxyResponse{StatusCode: 200, Body: string(json)}, nil
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

