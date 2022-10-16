package initializers

import "os"

var mongoURI string

func LoadEnvVars(){
	mongoURI = os.Getenv("MONGODB_URI")
}