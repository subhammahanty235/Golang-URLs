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

// decoding the json data

func decodeJson() {
	jsondatasample := []byte(`
	{
		"Name": "subham",
		"Email": "subham@gmail.com",
		"Password": "subham@12345"

	}
	`)

	var user creds

	isvalidJson := json.Valid(jsondatasample)
	if isvalidJson {
		fmt.Println("json is valid")
		json.Unmarshal(jsondatasample, &user)

		fmt.Printf("%#v\n", user)
	} else {
		fmt.Println("Json is not valid")
	}

	// json using map instead of structure

	var myuser map[string]interface{}
	json.Unmarshal(jsondatasample, &myuser)
	fmt.Printf("%#v\n", myuser)

	// printing data using loop

	for key, value := range myuser {
		fmt.Printf(" Key : %v and value is %v \n ", key, value)
	}

}

func main() {
	// createJson()
	decodeJson()
}
