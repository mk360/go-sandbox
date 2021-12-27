package main

import (
	"encoding/json"
	"fmt"
	"hello/hellofunc"
	"io/ioutil"
	"log"
	"net/http"
)

type S struct {
	X int
}

type Dummy struct {
	UserId    int64  `json:"userId"`
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	hellofunc.Hello("fire emblem")
	var d Dummy = Dummy{}
	readThing("https://jsonplaceholder.typicode.com/todos/1", d)
}

func readThing(url string, s map[string]int) {
	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(body, &s)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(s)
	}
}
