package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiurl = "https://jsonplaceholder.typicode.com/todos"

func main() {

	responce, err := http.Get(apiurl)

	if err != nil {
		panic(err)
	}
	defer responce.Body.Close()
	fmt.Printf("%T\n ", responce)
	// fmt.Println(responce)
	databytes, err := ioutil.ReadAll(responce.Body)

	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content[1])

	// URL parsing

	res, err := url.Parse(apiurl)
	if err != nil {
		panic(err)
	}

	fmt.Println(res.RawQuery)

}
