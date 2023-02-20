package main

import (
	"encoding/json"
	"fmt"
)

type creds struct {
	Name     string
	Email    string
	Password string
}

func createJson() {
	users := []creds{
		{Name: "subham", Email: "subham@gmail.com", Password: "subham@1234"},
		{Name: "subham2", Email: "subham2@gmail.com", Password: "subham2@1234"},
		{Name: "subham3", Email: "subham3@gmail.com", Password: "subham3@1234"},
	}

	jsondata, err := json.Marshal(users)
	organisedjsondata, err := json.MarshalIndent(users, "", "\a")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", string(jsondata))
	fmt.Printf("%s \n", string(organisedjsondata))
}

func main() {
	createJson()
}
