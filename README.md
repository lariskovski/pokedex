# Pokedex Go API 

Project was made to learn Go, a recent but strong passion! I used another repository to a web page to get Pokemon data and serves it with this Go CRUD Pokemons API. The deployment is made using Serverless Framework which is very straighforward and has a simple interface.

## Requirements

- [Serverless](https://www.serverless.com/framework/docs/install-standalone)

- [Go 1.19](https://go.dev/dl/)

- AWS Credentials

- [MongoDB Cluster](https://www.mongodb.com/es/cloud)

- [upx](https://github.com/upx/upx/releases) (Optional)

## Deployment

Edit accordingly and export the needed env vars:

~~~~
$ export AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE
$ export AWS_SECRET_ACCESS_KEY=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
$ export AWS_DEFAULT_REGION=us-east-1
$ export MONGODB_URI="mongodb+srv://usr:pwd@cluster0.wasd.mongo.net"
~~~~

Build the code optionally using upx (make shirnk) to reduce package size then deploy with Serverless:
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

## Sources

- [Tech With Tim tutorial](https://www.youtube.com/watch?v=bj77B59nkTQ)

- [Creating a Go API using an ORM Tutorial](https://www.youtube.com/watch?v=VAGodyl84OY)

- [SHRINK YOUR GO BINARIES WITH THIS ONE WEIRD TRICK](https://words.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/)

- [Auth0 custom authorizers](https://auth0.com/docs/customize/integrations/aws/aws-api-gateway-custom-authorizers#create-the-custom-authorizers)

- Many more...