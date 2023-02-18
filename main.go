package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const apiurl = "https://jsonplaceholder.typicode.com/todos"

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

	fmt.Println(" Press 1 to Login")
	var num int
	fmt.Scan(&num)
	if num == 1 {
		perform_post_req()
	} else {
		fmt.Println("Thanks for using")
	}

}

func perform_post_req() {
	const myurl = "https://codersbook-server-app-production.up.railway.app/api/auth/login"
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

	fmt.Println(string(content))
}
