package main

import (
	"fmt"
	"os"

	"github.com/mdb/seaweed"
)

// passed in via Makefile
var version string

func main() {
	_ = seaweed.NewClient(os.Getenv("MAGIC_SEAWEED_API_KEY"))

	fmt.Println(version)
}
