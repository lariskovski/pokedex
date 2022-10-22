package main

import (
	"encoding/json"
	"log"

	db "github.com/lariskovski/pokedex/api/commons"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	// uuid "github.com/satori/go.uuid"
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

func init(){
	db.Connect()
}

func main() {
	lambda.Start(postPokemon)
}

func postPokemon(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	body := request.Body
	if !(body == "" || body == "{}") {
		// Transforms request body json into Pokemon struct
		var pokemon Pokemon
		json.Unmarshal([]byte(request.Body), &pokemon)
		// pokemon.Id = uuid.NewV4().String()
	
		_, err := db.Collection.InsertOne(db.Context, pokemon)
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

