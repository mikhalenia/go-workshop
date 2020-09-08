package main

import (
	"fmt"
	"os"
)

func main() {
	someKey := os.Getenv("TEST_ENV_KEY")
	fmt.Printf("Some key : %v\n", someKey)
}