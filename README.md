# Pokemons Go API 

Project was made to refresh my Python web scrapping skills and learn Go. It scraps a web page to get Pokemon data and serves it with Go CRUD Pokemons API. Go base resource: [Tech With Tim tutorial](https://www.youtube.com/watch?v=bj77B59nkTQ)

## Locally Running the API

`` go run .``

### GET Request

~~~~
$ curl localhost:8080/pokemons
[
    {
        "id": "#001",
        "name": "Bulbasaur",
        "types": [
            "grass",
            "poison"
        ],
        "image": "/swordshield/pokemon/small/001.png",
        "ability": "Overgrow Chlorophyll",
        "baseStats": {
            "attack": "49",
            "defense": "49",
            "healthPoints": "45",
            "specialDefense": "45",
            "speedAttack": "65",
            "speedDefense": "65"
        }
    },
    {
        "id": "#002",
        "name": "Ivysaur",
        "types": [
            "grass",
            "poison"
        ],
        "image": "/swordshield/pokemon/small/002.png",
        "ability": "Overgrow Chlorophyll",
        "baseStats": {
            "attack": "62",
            "defense": "63",
            "healthPoints": "60",
            "specialDefense": "60",
            "speedAttack": "80",
            "speedDefense": "80"
        }
    },
    {
        "id": "#007",
        "name": "Squirtle",
        "types": [
            "water"
        ],
        "image": "/swordshield/pokemon/small/007.png",
        "ability": "Torrent Rain Dish",
        "baseStats": {
            "attack": "48",
            "defense": "65",
            "healthPoints": "44",
            "specialDefense": "43",
            "speedAttack": "50",
            "speedDefense": "64"
        }
    }
~~~~

### POST Request

~~~~
$ curl -X POST localhost:8080/pokemons --include --header "Content-Type: application/json" -d @post.json
HTTP/1.1 201 Created
Content-Type: application/json; charset=utf-8
Date: Sat, 15 Oct 2022 21:33:02 GMT
Content-Length: 364

{
    "id": "#007",
    "name": "Squirtle",
    "types": [
        "water"
    ],
    "image": "/swordshield/pokemon/small/007.png",
    "ability": "Torrent Rain Dish",
    "baseStats": {
        "attack": "48",
        "defense": "65",
        "healthPoints": "44",
        "specialDefense": "43",
        "speedAttack": "50",
        "speedDefense": "64"
    }
~~~~