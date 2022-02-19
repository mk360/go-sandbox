package main

import "fmt"

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
	var i int = 0
	for {
		fmt.Println(i)
		fibonacci(i)
		i++
	}
}

func fibonacci(n int) int {
	if n >= 2 {
		return fibonacci(n-1) + fibonacci(n-2)
	} else {
		return 1
	}
}
