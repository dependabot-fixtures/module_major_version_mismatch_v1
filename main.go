package main

import (
	"encoding/json"
	"fmt"

	"github.com/alecthomas/jsonschema"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	reflector := &jsonschema.Reflector{}
	schema := reflector.Reflect(&User{})

	schemaJSON, err := json.MarshalIndent(schema, "", "  ")
	if err != nil {
		fmt.Println("Error generating JSON schema:", err)
		return
	}

	fmt.Println(string(schemaJSON))
}
