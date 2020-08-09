package server

import (
	"fmt"
)

func parseCommand(command string, args interface{}) (string, error) {
	if command != "foo" {
		return "", fmt.Errorf("invalid command %s", command)
	}

	return foo(args), nil
}

func foo(data interface{}) string {
	result := "foo "
	result += fmt.Sprintf("%v,", data)
	result = result[:len(result)-1] // Remove last comma
	return result
}
