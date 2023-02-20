package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const apiurl = "https://jsonplaceholder.typicode.com/users"

func main() {

	// responce, err := http.Get(apiurl)

	// if err != nil {
	// 	panic(err)
	// }
	// defer responce.Body.Close()
	// fmt.Printf("%T\n ", responce)
	// // fmt.Println(responce)
	// databytes, err := ioutil.ReadAll(responce.Body)

	// if err != nil {
	// 	panic(err)
	// }

	// content := string(databytes)

	// fmt.Println(content[1])

	// // URL parsing

	// res, err := url.Parse(apiurl)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(res.RawQuery)

	fmt.Println(" Press 1 for get request and 2 for post , 3 for json")

	var num int
	fmt.Scan(&num)
	if num == 2 {
		perform_post_req()
	} else if num == 1 {
		perfromgetreq()
	} else {
		fmt.Println("Thanks for Using")
	}

}
func perfromgetreq() {
	const myurl = apiurl

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}

func perform_post_req() {
	const myurl = "http://localhost:5000/api/auth/login"
	requestbody := strings.NewReader(`
		{
			"email":"test1@gmail.com",
			"password":"subham@1234",
		}
	`)

	response, err := http.Post(myurl, "applicaton/json", requestbody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	// fmt.Printf("TYpe of %T \n", requestbody)
	fmt.Println(string(content))
	// fmt.Println()
}
