package db

import (
	"fmt"
	"os"
)

func printURI() {
	uri := os.Getenv("MONGODB_URI")
	fmt.Println(uri)
}
