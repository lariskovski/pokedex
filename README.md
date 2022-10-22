# Pokemons Go API 

Project was made to refresh my Python web scrapping skills and learn Go. It scraps a web page to get Pokemon data and serves it with Go CRUD Pokemons API.

## Requirements

- [Serverless](https://www.serverless.com/framework/docs/install-standalone)

- [Go 1.19](https://go.dev/dl/)

- AWS Credentials

- [MongoDB](https://www.mongodb.com/es/cloud)

- [upx](https://github.com/upx/upx/releases) (Optional)

## Deployment

Export the needed env vars:

~~~~
$ export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
$ export AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
$ export AWS_DEFAULT_REGION=us-west-2
$ export MONGODB_URI="mongodb+srv://usr:pwd@cluster0.wasd.mongo.net"
~~~~

Build the code optionally using upx (make shirnk) to reduce package size then deploy:
~~~~
$ make build
$ make shrink
$ make deploy
~~~~

Cleaning everything local and remote (be careful):

~~~~
make clean
make remove
~~~~

## Custom Authorizer

By default API Gateway deployments aren't secure. Optionally you may add a custom authorizer following the instructions on [this](https://github.com/lariskovski/jwt-rsa-aws-custom-authorizer-serverless) repo then uncommenting the authorizer part on `serverless.yml` and exporting the `LAMBDA_AUTHORIZER_ARN`environment variable.

## Requests Examples

- POST:
~~~~
$ curl -X POST localhost:8080/pokemons --include --header "Content-Type: application/json" -d @post.json
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Sun, 16 Oct 2022 19:34:40 GMT
Content-Length: 48

{
    "InsertedID": "634c5cce0099c492e89f51aa"
}
~~~~

- GET:

~~~~
$ curl localhost:8080/pokemons
[
    {
        "_id": "634c5917f411a54fd21e6f7b",
        "ability": "Torrent Rain Dish",
        "baseStats": {
            "attack": "48",
            "defense": "65",
            "healthPoints": "100",
            "specialDefense": "43",
            "speedAttack": "50",
            "speedDefense": "64"
        },
        "basestats": {
            "attack": "48",
            "defense": "65",
            "healthPoints": "44",
            "specialDefense": "43",
            "speedAttack": "50",
            "speedDefense": "64"
        },
        "image": "/swordshield/pokemon/small/007.png",
        "name": "Squirtlexxxx",
        "types": [
            "water"
        ]
    },
    {
        "_id": "634c5cce0099c492e89f51aa",
        "ability": "Torrent Rain Dish",
        "basestats": {
            "attack": "48",
            "defense": "65",
            "healthPoints": "44",
            "specialDefense": "43",
            "speedAttack": "50",
            "speedDefense": "64"
        },
        "image": "/swordshield/pokemon/small/007.png",
        "name": "Squirtle",
        "types": [
            "water"
        ]
    }
]
~~~~

- PUT:

~~~~
$ curl -X PUT localhost:8080/pokemons/634c5cce0099c492e89f51aa -d @put.json
{
    "message": "Pokemon updated."
}
~~~~

- DELETE:

~~~~
$ curl -X DELETE localhost:8080/pokemons/634c5cce0099c492e89f51aa
{
    "message": "Pokemon deleted."
}
~~~~

## Sources

- [Tech With Tim tutorial](https://www.youtube.com/watch?v=bj77B59nkTQ)

- [Creating a Go API using an ORM Tutorial](https://www.youtube.com/watch?v=VAGodyl84OY)

- [SHRINK YOUR GO BINARIES WITH THIS ONE WEIRD TRICK](https://words.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/)

- Many (MANY) more...