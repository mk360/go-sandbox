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
	UserId    float32 `json:"userId"`
	Id        int64   `json:"A"`
	Title     string  `json:"title"`
	Completed bool    `json:"completed"`
}

func main() {
	hellofunc.Hello("fire emblem")
	var d Dummy = Dummy{}
	readThing("http://localhost:3000", &d)
	// this quickly-cooked nodejs api server returns {
	//	A: 3.94
	// }
}

func readThing(url string, s *Dummy) {
	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)
	err := json.Unmarshal(body, &s)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(s)
	}
}
