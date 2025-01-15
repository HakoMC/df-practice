package db

import (
	"os"
)

func printURI() string {
	uri := os.Getenv("MONGODB_URI")
	return uri
}
