package main

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"go.mongodb.org/mongo-driver/bson"
	db "github.com/lariskovski/pokedex/api/commons"
)

type Response struct {
	Message string `json:"message"`
}

func init(){
	db.Connect()
}

func main() {
	lambda.Start(deletePokemon)
}


func deletePokemon(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	name := request.PathParameters["name"]
	if name != "" {
		deleteResult, err := db.Collection.DeleteOne(db.Context, bson.D{{Key: "name", Value: name }})
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

